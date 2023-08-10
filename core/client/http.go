package client

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"zrDispatch/common/utils"
	"zrDispatch/core/config"
	"zrDispatch/core/slog"

	"go.uber.org/zap"
)

type ResponseJson struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}
type CountJson struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data int    `json:"data"`
}

type HostInfoJson struct {
	Cpuinfos  map[string]interface{} `json:"cpuinfos"`
	Hostinfos map[string]interface{} `json:"hostinfos"`
	Meminfos  map[string]interface{} `json:"meminfos"`
	Netinfos  []interface{}          `json:"netinfos"`
	Netspeed  []interface{}          `json:"netspeed"`
	Parts     []interface{}          `json:"parts"`
}

func getUrl(ip string, port int, path string) string {
	url := "https://" + ip + ":" + utils.GetInterfaceToString(port) + path

	//log.Info("url :" + url)

	return url
}

func GetCli(timeout time.Duration) (*http.Client, string) {

	// setup a http client
	httpTransport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	httpClient := &http.Client{Transport: httpTransport, Timeout: timeout}

	// set our socks5 as the dialer
	// create a socks5 dialer
	// num := utils.RanNum(len(define.ProxyMap))
	// addr := define.ProxyMap[num]
	// // slog.Println(slog.WARN, "addr", addr)
	// if addr != "0" {
	// 	dialer, err := proxy.SOCKS5("tcp", addr, nil, proxy.Direct)
	// 	if err != nil {
	// 		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
	// 		os.Exit(1)
	// 	}
	// 	httpTransport.Dial = dialer.Dial
	// }

	return httpClient, ""
}

func Send(url, data string) (body []byte, err error) {
	req, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(data))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", config.CoreConf.BasicAuth)

	cli, _ := GetCli(20 * time.Second)
	resp, err := cli.Do(req)
	if err != nil {
		slog.Println(slog.DEBUG, zap.Error(err))
		return
	}
	body, err = ioutil.ReadAll(resp.Body)

	return
}

func RequestHeartBeat(ip string, port int) map[string]interface{} {
	url := getUrl(ip, port, "/v1/check/heartbeat")

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Add("Authorization", config.CoreConf.BasicAuth)

	var responseJson ResponseJson

	cli, addr := GetCli(10 * time.Second)
	resp, err := cli.Do(req)

	if err != nil {
		slog.Println(slog.DEBUG, "发送心跳失败", zap.Error(err))
		return nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	//{"code":200,"data":{"runningTasks":"3412341234","time":"1672733119","version":"1.1.1"},"msg":""}

	// slog.Println(slog.DEBUG, ip, "=====", string(body))
	if err != nil {
		slog.Println(slog.DEBUG, "读取response失败", addr, zap.Error(err))
		return nil
	}
	if err = json.Unmarshal(body, &responseJson); err != nil {
		slog.Println(slog.DEBUG, "json读取失败", addr, zap.Error(err))
		return nil
	}

	if responseJson.Code != 200 {
		slog.Println(slog.DEBUG, "心跳返回错误", addr, zap.Error(err))
		return nil
	}

	return responseJson.Data
}
