package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	"zrDispatch/common/log"
	"zrDispatch/core/config"
	"zrDispatch/core/utils/define"

	"go.uber.org/zap"
)

func HostInfo(hostInfo *define.Host) HostInfoJson {
	req, _ := http.NewRequest(http.MethodGet, getUrl(hostInfo.Ip, hostInfo.ServicePort, "/v1/check/hostinfo"), nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", config.CoreConf.BasicAuth)

	var responseJson HostInfoJson

	cli, _ := GetCli(20 * time.Second)
	resp, err := cli.Do(req)

	if err != nil {
		log.Debug("发送心跳失败", zap.Error(err))
	}
	body, err := ioutil.ReadAll(resp.Body)
	//{"code":200,"data":{"runningTasks":"3412341234","time":"1672733119","version":"1.1.1"},"msg":""}
	if err != nil {
		log.Error("读取response失败", zap.Error(err))
	}
	if err = json.Unmarshal(body, &responseJson); err != nil {
		log.Error("json读取失败", zap.Error(err))
	}

	return responseJson
}
