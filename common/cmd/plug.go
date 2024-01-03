package cmd

import (
	"os"
	"os/exec"
	"strings"
	"zrDispatch/common/utils"
	"zrDispatch/core/slog"
)

func Plug(taskId, ip, plugname, cmd string) (string, error) {

	if cmd == "" {
		return "", nil
	}

	ipstr := strings.Join(utils.GetIpArr(ip), " ")

	dir, _ := os.Getwd()

	path := dir + "/plugres/" + plugname

	if !utils.PathExist(path) {
		err := os.MkdirAll(path, 0777)
		if err != nil {
			slog.Println(slog.WARN, "新建日志目录失败")
		}
	}

	res := path + "/" + taskId + ".res"

	if strings.Contains(cmd, "nmap") {
		if !utils.PathExist("/usr/bin/nmap") {
			slog.Println(slog.DEBUG, "nmap不存在")
			return "", nil
		}
	} else if strings.Contains(cmd, "python") {
		cmd = "cd " + dir + " && " + cmd
	} else {
		exe := strings.Split(cmd, " ")[0]
		exe = dir + "/" + exe
		if !utils.PathExist(exe) {
			slog.Println(slog.DEBUG, "文件不存在", exe)
			return "", nil
		}
		cmd = dir + "/" + cmd
		ipstr = strings.Join(utils.GetIpArr(ip), ",")
	}

	if strings.Contains(cmd, "{res}") {
		cmd = strings.Replace(cmd, "{res}", res, -1)
	} else if strings.Contains(cmd, "{SingleIP}") {
		cmd = cmd + " >>" + res
	} else {
		cmd = cmd + " >" + res
	}

	if strings.Contains(cmd, "{SingleIP}") {
		ipArr := utils.GetIpArr(ip)
		for _, sip := range ipArr {
			utils.WriteAppend(res, "ip:"+sip+"   ")
			cmdip := strings.Replace(cmd, "{SingleIP}", sip, -1)
			slog.Println(slog.DEBUG, cmdip)

			cmd1 := exec.Command("bash", "-c", cmdip)

			cmdres, err := cmd1.CombinedOutput()

			slog.Println(slog.DEBUG, string(cmdres), "==========", err)
			utils.WriteAppendHh(res, string(cmdres))
		}

	} else {
		cmd = strings.Replace(cmd, "{ip}", ipstr, -1)
		slog.Println(slog.DEBUG, cmd)

		cmd1 := exec.Command("bash", "-c", cmd)

		_, err := cmd1.CombinedOutput()
		if err != nil {
			slog.Println(slog.ERROR, err)

			return "", err

		}
		return res, nil
	}

	return res, nil
}

func CheckExec(cmdstr, filename string) (err error) {
	dir, _ := os.Getwd()

	Exec("chmod +x " + dir + "/" + filename)

	cmdstr = dir + "/" + cmdstr

	cmdstr = strings.Replace(cmdstr, "{ip}", "127.0.0.1", -1)

	cmdstr = strings.Replace(cmdstr, "{res}", "test.res", -1)

	_, err = Exec(cmdstr)

	return
}

func CheckScript(cmdstr, filename string) (err error) {
	dir, _ := os.Getwd()

	Exec("chmod +x " + dir + "/" + filename)

	cmdstr = "cd " + dir + " && " + cmdstr

	cmdstr = strings.Replace(cmdstr, "{ip}", "127.0.0.1", -1)

	cmdstr = strings.Replace(cmdstr, "{SingleIP}", "127.0.0.1", -1)

	cmdstr = strings.Replace(cmdstr, "{res}", "test.res", -1)

	_, err = Exec(cmdstr)

	return
}
