package client

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"strings"
	"zrDispatch/common/utils"
	"zrDispatch/core/slog"
	"zrDispatch/core/utils/define"
	"zrDispatch/models"

	"github.com/bytedance/sonic"
	"go.uber.org/zap"
)

/**
{
    // "service_type": "probe",
    "task_id": "11111",
    "addrs": [
        "127.0.0.1:80",
        "127.0.0.1:3306"
    ],
    "payload": [
        {
            "payload": "R0VUIC8xLzIvMy80LzUgSFRUUC8xLjFcclxuVXNlci1BZ2VudDogTW96aWxsYS81LjBcclxuQWNjZXB0OiB0ZXh0L2h0bWwsYXBwbGljYXRpb24veGh0bWwreG1sLGFwcGxpY2F0aW9uL3htbDtxPTAuOSxpbWFnZS93ZWJwLCovKjtxPTAuOFxyXG5BY2NlcHQtRW5jb2Rpbmc6IGd6aXAsIGRlZmxhdGVcclxuQ29ubmVjdGlvbjogY2xvc2U=",
            "probe_name": "aaaaa"
        }
    ],
    "timeout": 5,
    "threads": 5
}
**/

type RequestData struct {
	TaskId  string   `json:"task_id"`
	Addrs   []string `json:"addrs"`
	Payload []define.Pyload
	Timeout int `json:"timeout"`
	Threads int `json:"threads"`
}

func RunTask(hostInfo *define.Host, taskdata *define.DetailTask) error {

	var reqD RequestData
	reqD.TaskId = taskdata.RunTaskId
	reqD.Timeout = taskdata.Timeout
	reqD.Threads = taskdata.Threads
	reqD.Addrs = utils.GetAddrs(taskdata.Ip, taskdata.Port)
	reqD.Payload, _ = models.GetPayload(taskdata.ProbeId)

	//base64的payload
	// slog.Println(slog.DEBUG, "===任务开始===", reqD.Payload)
	jsonData, _ := sonic.Marshal(&reqD)

	// slog.Println(slog.DEBUG, string(jsonData))

	var responseJson ResponseJson

	url := getUrl(hostInfo.Ip, hostInfo.ServicePort, "/v1/"+taskdata.TaskType.Value()+"/start")

	body, err := Send(url, string(jsonData))

	if err != nil {
		slog.Println(slog.DEBUG, "err", zap.Error(err))
		return err
	}

	if err = json.Unmarshal(body, &responseJson); err != nil {
		slog.Println(slog.DEBUG, "json读取失败==", zap.Error(err))
		return err
	}

	if responseJson.Code != 200 {
		slog.Println(slog.DEBUG, "===任务返回错误:"+responseJson.Msg)
		return err
	}
	return nil
}

type GetTaskPressRequestData struct {
	TaskId string `json:"task_id"`
}

type GetTaskPressResponsetData struct {
	All  int    `json:"all_addr"`
	Ok   int    `json:"ok_addr"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 任务是否完成
func GetTaskPress(hostInfo *define.Host, taskdata *define.DetailTask) (bool, error) {

	url := getUrl(hostInfo.Ip, hostInfo.ServicePort, "/v1/"+taskdata.TaskType.Value()+"/progress")

	var reqD GetTaskPressRequestData
	reqD.TaskId = taskdata.RunTaskId

	jsonData, _ := sonic.Marshal(&reqD)

	body, err := Send(url, string(jsonData))

	var responseJson GetTaskPressResponsetData

	slog.Println(slog.DEBUG, string(body))
	if err = json.Unmarshal(body, &responseJson); err != nil {
		slog.Println(slog.DEBUG, "json读取失败==", zap.Error(err))
		return false, nil
	}

	slog.Println(slog.DEBUG, hostInfo.Ip, "==", taskdata.RunTaskId, "==", responseJson.Code, responseJson.All, responseJson.Ok)

	if responseJson.Code != 200 {
		return false, errors.New(responseJson.Msg)
	}

	if responseJson.Code == 200 && responseJson.All != 0 && responseJson.All == responseJson.Ok {
		return true, nil
	} else {
		models.UpdatePress(taskdata.RunTaskId, responseJson.Ok*100/responseJson.All)
	}

	return false, nil
}

func GetTaskRes(hostInfo *define.Host, taskdata *define.DetailTask) error {

	// slog.Println(slog.DEBUG, "GetTaskRes")

	url := getUrl(hostInfo.Ip, hostInfo.ServicePort, "/v1/"+taskdata.TaskType.Value()+"/result")

	var reqD GetTaskPressRequestData
	reqD.TaskId = taskdata.RunTaskId

	jsonData, _ := sonic.Marshal(&reqD)

	body, err := Send(url, string(jsonData))

	if err != nil {
		slog.Println(slog.DEBUG, "GetTaskRes", zap.Error(err))
		return err
	}

	slog.Println(slog.DEBUG, hostInfo.Ip, "====", string(body))

	var responseJson *define.TaskRes

	if err = json.Unmarshal(body, &responseJson); err != nil {
		slog.Println(slog.DEBUG, "json读取失败==", zap.Error(err))
		return err
	}

	// fmt.Println("%#", responseJson)

	//探针的结果
	if responseJson.Code == 200 && len(responseJson.Data) > 0 {
		for _, obj := range responseJson.Data {
			if obj.ProbeResult.ResPlain != "" {

				var ps define.ProbeResCreate
				ps.Ctime = utils.GetTimeStr()

				Arr := strings.Split(obj.ProbeResult.ReqInfo.Addr, ":")
				ps.IP = Arr[0]
				ps.Port = Arr[1]
				ps.Pname = obj.ProbeResult.ReqInfo.ProbeName
				ps.Hex = obj.ProbeResult.ResHex
				ps.Res = obj.ProbeResult.ResPlain
				ps.RunTaskID = taskdata.RunTaskId

				models.AddProbeRes(ps)
			}

		}
	}
	//证书的结果
	if responseJson.Code == 200 && len(responseJson.Data) > 0 {
		for _, obj := range responseJson.Data {
			if obj.SslResult.Cert.CertBase64 != "" {

				Arr := strings.Split(obj.ProbeResult.ReqInfo.Addr, ":")
				obj.SslResult.Cert.Ip = Arr[0]
				obj.SslResult.Cert.Port = Arr[1]
				obj.SslResult.Cert.Probe_name = obj.ProbeResult.ReqInfo.ProbeName
				obj.SslResult.Cert.RunTaskID = taskdata.RunTaskId
				models.AddCertRes(obj.SslResult.Cert)
			}

		}
	}

	//端口的结果
	if responseJson.Code == 200 && len(responseJson.Res) > 0 {
		for _, str := range responseJson.Res {
			var portInfo define.PortScan
			if err = json.Unmarshal([]byte(str), &portInfo); err != nil {
				slog.Println(slog.DEBUG, "json读取失败==", zap.Error(err))
				return err
			}

			portInfo.Hex = hex.Dump([]byte(portInfo.Response))

			portInfo.Ctime = utils.GetTimeStr()
			models.AddPortRes(portInfo)
		}
	}

	return nil
}
