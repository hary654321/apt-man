// Package sshclient implements an SSH client.
package sshclient

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
	"zrDispatch/common/utils"
	"zrDispatch/core/redis"
	"zrDispatch/core/slog"
	"zrDispatch/core/utils/define"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"golang.org/x/net/proxy"
)

type remoteScriptType byte
type remoteShellType byte

const (
	cmdLine remoteScriptType = iota
	rawScript
	scriptFile

	interactiveShell remoteShellType = iota
	nonInteractiveShell
)

// A Client implements an SSH client that supports running commands and scripts remotely.
type Client struct {
	client *ssh.Client
}

// DialWithPasswd starts a client connection to the given SSH server with passwd authmethod.
func DialWithPasswd(addr, user, passwd string) (*Client, error) {

	slog.Println(slog.DEBUG, addr, user, passwd)
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(passwd),
		},
		HostKeyCallback: ssh.HostKeyCallback(func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil }),
	}

	return Dial("tcp", addr, config)

}

func proxiedSSHClient(proxyAddress, sshServerAddress string, sshConfig *ssh.ClientConfig) (*ssh.Client, error) {

	dialer, err := proxy.SOCKS5("tcp", proxyAddress, nil, proxy.Direct)
	if err != nil {
		return nil, err
	}

	conn, err := dialer.Dial("tcp", sshServerAddress)
	if err != nil {
		return nil, err
	}

	c, chans, reqs, err := ssh.NewClientConn(conn, sshServerAddress, sshConfig)
	if err != nil {
		return nil, err
	}

	return ssh.NewClient(c, chans, reqs), nil
}

func netConn(addr string) (net.Conn, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to connect to %v", err))
	}
	return conn, nil
}

func Dialer(conf *ssh.ClientConfig, c net.Conn, host string) (*ssh.Client, error) {
	conn, chans, reqs, err := ssh.NewClientConn(c, host, conf)
	if err != nil {
		return nil, err
	}
	return ssh.NewClient(conn, chans, reqs), nil
}

// DialWithKey starts a client connection to the given SSH server with key authmethod.
func DialWithKey(addr, user, keyfile string) (*Client, error) {
	key, err := ioutil.ReadFile(keyfile)
	if err != nil {
		return nil, err
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.HostKeyCallback(func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil }),
	}

	return Dial("tcp", addr, config)
}

// DialWithKeyWithPassphrase same as DialWithKey but with a passphrase to decrypt the private key
func DialWithKeyWithPassphrase(addr, user, keyfile string, passphrase string) (*Client, error) {
	key, err := ioutil.ReadFile(keyfile)
	if err != nil {
		return nil, err
	}

	signer, err := ssh.ParsePrivateKeyWithPassphrase(key, []byte(passphrase))
	if err != nil {
		return nil, err
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.HostKeyCallback(func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil }),
	}

	return Dial("tcp", addr, config)
}

// Dial starts a client connection to the given SSH server.
// This wraps ssh.Dial.
func Dial(network, addr string, config *ssh.ClientConfig) (*Client, error) {
	client, err := ssh.Dial(network, addr, config)
	if err != nil {
		return nil, err
	}
	return &Client{
		client: client,
	}, nil
}

// Close closes the underlying client network connection.
func (c *Client) Close() error {
	return c.client.Close()
}

// UnderlyingClient get the underlying client.
func (c *Client) UnderlyingClient() *ssh.Client {
	return c.client
}

// Dial initiates a Client to the addr from the remote host.
func (c *Client) Dial(network, addr string, config *ssh.ClientConfig) (*Client, error) {
	conn, err := c.client.Dial(network, addr)
	if err != nil {
		return nil, err
	}

	sshConn, chans, reqs, err := ssh.NewClientConn(conn, addr, config)
	if err != nil {
		return nil, err
	}

	client := ssh.NewClient(sshConn, chans, reqs)

	return &Client{client: client}, nil
}

// Cmd creates a RemoteScript that can run the command on the client. The cmd string is split on newlines and each line is executed separately.
func (c *Client) Cmd(cmd string) *RemoteScript {
	//slog.Println(slog.DEBUG, c)
	if c == nil {
		return nil
	}
	return &RemoteScript{
		_type:  cmdLine,
		client: c.client,
		script: bytes.NewBufferString(cmd + "\n"),
	}
}

// Upload a local file to remote server!
func (c *Client) Upload(localPath string, remotePath string) (err error) {

	local, err := os.Open(localPath)
	if err != nil {
		slog.Println(slog.DEBUG, err)
		return
	}
	defer local.Close()

	ftp, err := c.NewSftp()
	if err != nil || ftp == nil {
		slog.Println(slog.DEBUG, err)
		return
	}
	defer ftp.Close()

	remote, err := ftp.Create(remotePath)
	if err != nil {
		slog.Println(slog.DEBUG, err)
		return
	}
	defer remote.Close()

	_, err = io.Copy(remote, local)
	return
}

func (c *Client) Download(srcPath, dstPath, name string) error {

	slog.Println(slog.DEBUG, srcPath, "==============", name)
	ftp, err := c.NewSftp()
	if err != nil {
		slog.Println(slog.DEBUG, err)
		return err
	}
	srcFile, err := ftp.Open(srcPath) //远程

	if err != nil {
		slog.Println(slog.DEBUG, err)
		return err
	}

	_, err1 := os.Stat(dstPath)
	if err1 != nil {
		os.MkdirAll(dstPath, 0777)
	}

	dstFile, _ := os.Create(dstPath + "/" + name) //本地
	defer func() {
		_ = srcFile.Close()
		_ = dstFile.Close()
	}()
	if _, err := srcFile.WriteTo(dstFile); err != nil {
		slog.Println(slog.DEBUG, "error occurred", err)
		return err
	}
	slog.Println(slog.DEBUG, "文件下载完毕")
	return nil
}

// NewSftp returns new sftp client and error if any.
func (c *Client) NewSftp(opts ...sftp.ClientOption) (*sftp.Client, error) {
	return sftp.NewClient(c.client, opts...)
}

// Script creates a RemoteScript that can run the script on the client.
func (c *Client) Script(script string) *RemoteScript {
	return &RemoteScript{
		_type:  rawScript,
		client: c.client,
		script: bytes.NewBufferString(script + "\n"),
	}
}

// ScriptFile creates a RemoteScript that can read a local script file and run it remotely on the client.
func (c *Client) ScriptFile(fname string) *RemoteScript {
	return &RemoteScript{
		_type:      scriptFile,
		client:     c.client,
		scriptFile: fname,
	}
}

// A RemoteScript represents script that can be run remotely.
type RemoteScript struct {
	client     *ssh.Client
	_type      remoteScriptType
	script     *bytes.Buffer
	scriptFile string
	err        error

	stdout io.Writer
	stderr io.Writer
}

// Run runs the script on the client.
//
// The returned error is nil if the command runs, has no problems
// copying stdin, stdout, and stderr, and exits with a zero exit
// status.
func (rs *RemoteScript) Run() error {
	if rs.err != nil {
		fmt.Println(rs.err)
		return rs.err
	}

	if rs._type == cmdLine {
		return rs.runCmds()
	} else if rs._type == rawScript {
		return rs.runScript()
	} else if rs._type == scriptFile {
		return rs.runScriptFile()
	} else {
		return errors.New("Not supported RemoteScript type")
	}
}

// Output runs the script on the client and returns its standard output.
func (rs *RemoteScript) Output() string {
	if rs == nil {
		return ""
	}
	if rs.stdout != nil {
		fmt.Println("Stdout already set")
	}
	var out bytes.Buffer
	rs.stdout = &out
	err := rs.Run()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(out.Bytes()))
	}

	return string(out.Bytes())
}

// SmartOutput runs the script on the client. On success, its standard ouput is returned. On error, its standard error is returned.
func (rs *RemoteScript) SmartOutput() ([]byte, error) {
	if rs.stdout != nil {
		return nil, errors.New("Stdout already set")
	}
	if rs.stderr != nil {
		return nil, errors.New("Stderr already set")
	}

	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)
	rs.stdout = &stdout
	rs.stderr = &stderr
	err := rs.Run()
	if err != nil {
		return stderr.Bytes(), err
	}
	return stdout.Bytes(), err
}

// Cmd appends a command to the RemoteScript.
func (rs *RemoteScript) Cmd(cmd string) *RemoteScript {
	_, err := rs.script.WriteString(cmd + "\n")
	if err != nil {
		rs.err = err
	}
	return rs
}

// SetStdio specifies where its standard output and error data will be written.
func (rs *RemoteScript) SetStdio(stdout, stderr io.Writer) *RemoteScript {
	rs.stdout = stdout
	rs.stderr = stderr
	return rs
}

func (rs *RemoteScript) runCmd(cmd string) error {
	session, err := rs.client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	session.Stdout = rs.stdout
	session.Stderr = rs.stderr

	if err := session.Run(cmd); err != nil {
		return err
	}
	return nil
}

func (rs *RemoteScript) runCmds() error {
	for {
		statment, err := rs.script.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if err := rs.runCmd(statment); err != nil {
			return err
		}
	}

	return nil
}

func (rs *RemoteScript) runScript() error {
	session, err := rs.client.NewSession()
	if err != nil {
		return err
	}

	session.Stdin = rs.script
	session.Stdout = rs.stdout
	session.Stderr = rs.stderr

	if err := session.Shell(); err != nil {
		return err
	}
	if err := session.Wait(); err != nil {
		return err
	}

	return nil
}

func (rs *RemoteScript) runScriptFile() error {
	var buffer bytes.Buffer
	file, err := os.Open(rs.scriptFile)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(&buffer, file)
	if err != nil {
		return err
	}

	rs.script = &buffer
	return rs.runScript()
}

// A RemoteShell represents a login shell on the client.
type RemoteShell struct {
	client         *ssh.Client
	requestPty     bool
	terminalConfig *TerminalConfig

	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
}

// A TerminalConfig represents the configuration for an interactive shell session.
type TerminalConfig struct {
	Term   string
	Height int
	Weight int
	Modes  ssh.TerminalModes
}

// Terminal create a interactive shell on client.
func (c *Client) Terminal(config *TerminalConfig) *RemoteShell {
	return &RemoteShell{
		client:         c.client,
		terminalConfig: config,
		requestPty:     true,
	}
}

// Shell create a noninteractive shell on client.
func (c *Client) Shell() *RemoteShell {
	return &RemoteShell{
		client:     c.client,
		requestPty: false,
	}
}

// SetStdio specifies where the its standard output and error data will be written.
func (rs *RemoteShell) SetStdio(stdin io.Reader, stdout, stderr io.Writer) *RemoteShell {
	rs.stdin = stdin
	rs.stdout = stdout
	rs.stderr = stderr
	return rs
}

// Start starts a remote shell on client.
func (rs *RemoteShell) Start() error {
	session, err := rs.client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	if rs.stdin == nil {
		session.Stdin = os.Stdin
	} else {
		session.Stdin = rs.stdin
	}
	if rs.stdout == nil {
		session.Stdout = os.Stdout
	} else {
		session.Stdout = rs.stdout
	}
	if rs.stderr == nil {
		session.Stderr = os.Stderr
	} else {
		session.Stderr = rs.stderr
	}

	if rs.requestPty {
		tc := rs.terminalConfig
		if tc == nil {
			tc = &TerminalConfig{
				Term:   "xterm",
				Height: 40,
				Weight: 80,
			}
		}
		if err := session.RequestPty(tc.Term, tc.Height, tc.Weight, tc.Modes); err != nil {
			return err
		}
	}

	if err := session.Shell(); err != nil {
		return err
	}

	if err := session.Wait(); err != nil {
		return err
	}

	return nil
}

func Exec(ctx context.Context, hostInfo *define.Host, cmd string, ch chan string) {
	client, err := DialWithPasswd(hostInfo.Ip+":"+utils.GetInterfaceToString(hostInfo.SshPort), hostInfo.SshUser, hostInfo.SshPwd)
	if err != nil {
		slog.Println(slog.WARN, err)
		return
	}
	//defer client.Close()

	res := client.Cmd(cmd).Output()
	clear := client.Cmd("ps -ef | grep root@notty | grep -v grep | awk '{print $2}' | xargs kill -9").Output()
	slog.Println(slog.DEBUG, hostInfo.Ip, "clear", clear, cmd)

	ch <- res
	select {
	case <-ctx.Done():
		close(ch)
		slog.Println(slog.DEBUG, ctx.Err())
		//client.Close()
		return
	default:
		//client.Close()
		//slog.Println(slog.DEBUG,n, "default")
	}
}

func ExecTimeOut(hostInfo *define.Host, cmd string) (string, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	ch := make(chan string)
	go Exec(ctx, hostInfo, cmd, ch)

	for {
		select {
		case v, ok := <-ch:
			slog.Println(slog.WARN, "执行完了", hostInfo.Ip, "----", cmd, v, ok)
			return v, ok
		case <-ctx.Done():
			slog.Println(slog.WARN, "超时了", hostInfo.Ip, "------", cmd)
			return "超时了", false
		}
	}
}
func Start(hostInfo *define.Host, force bool) (string, bool) {

	cli := redis.GetClient()

	timek := hostInfo.Ip + "restart"
	tc, _ := cli.Get(timek).Int64()

	nowtime, _ := strconv.ParseInt(utils.GetTime(), 10, 64)

	if nowtime-tc < 60*5 && !force {
		slog.Println(slog.WARN, hostInfo.Ip, "正在重启中。。。。", tc)
		return "正在重启中。。。", false
	}
	cli.Set(timek, utils.GetTime(), 20*60&time.Second)

	restartTime := hostInfo.Ip + "restartTime"

	s2, res2 := ExecTimeOut(hostInfo, "cd /zrtx/apt  && ./apt-scan >> /zrtx/log/cyberspace/worker.log &")

	cli.Incr(restartTime)

	return s2, res2
}

func Restart(hostInfo *define.Host) (string, bool) {

	return "", true
}

func ServiceLog(hostInfo *define.Host) (string, bool) {

	return ExecTimeOut(hostInfo, "tail -10 /zrtx/log/cyberspace/apt-scan.log")

}

func CleanRes(hostInfo *define.Host) (string, bool) {

	return ExecTimeOut(hostInfo, "find /zrtx/log/cyberspace  -mtime +1 -name \"*\" | xargs -I {} rm -rf {}")

}

func ServiceCmd(hostInfo *define.Host) (string, bool) {

	return ExecTimeOut(hostInfo, "ps -ef | grep ./apt")

}

func ResLog(hostInfo *define.Host) (string, bool) {

	return ExecTimeOut(hostInfo, "tail -10 /zrtx/log/cyberspace/ipInfo"+utils.GetHour()+".json")

}

func ResCount(hostInfo *define.Host, date string) (int, bool) {

	res, v := ExecTimeOut(hostInfo, "wc -l /zrtx/log/cyberspace/ipInfo"+date+".json")

	resArr := strings.Split(res, " ")

	num, _ := strconv.Atoi(resArr[0])

	return num, v

}

func ResIP(hostInfo *define.Host, date string) (int, bool) {

	res, v := ExecTimeOut(hostInfo, "wc -l /zrtx/log/cyberspace/ipLast"+date+".txt")

	resArr := strings.Split(res, " ")

	num, _ := strconv.Atoi(resArr[0])

	return num, v

}

func ResSize(hostInfo *define.Host, date string) (string, bool) {

	res, v := ExecTimeOut(hostInfo, " du -sh /zrtx/log/cyberspace/ipInfo"+date+".json")

	resArr := strings.Split(res, "/")

	return resArr[0], v
}

func GetDomain(hostInfo *define.Host, date string) (string, bool) {

	res, v := ExecTimeOut(hostInfo, " tail -1000 /zrtx/log/cyberspace/domain"+date+".json")

	return res, v
}

func CleanLog(hostInfo *define.Host) (string, bool) {

	res, v := ExecTimeOut(hostInfo, "find /zrtx/log/cyberspace  -mtime +1 -name \"*\" | xargs -I {} rm -rf {} && echo '' > /zrtx/log/cyberspace/worker.log")

	return res, v
}
