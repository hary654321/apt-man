package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
	"zrDispatch/common/utils"
	"zrDispatch/core/config"
	"zrDispatch/core/slog"

	"go.uber.org/zap"
)

type LoginRes struct {
	ReturnCode int `json:"returnCode"`
	Data       struct {
		Clients   []interface{} `json:"clients"`
		LoginName string        `json:"loginName"`
		Name      string        `json:"name"`
		InfoMap   struct {
			Token        string `json:"token"`
			RefreshToken string `json:"refreshToken"`
		} `json:"infoMap"`
		UserID string `json:"userId"`
	} `json:"data"`
	ErrorMsg string `json:"errorMsg"`
}

func getSsoUrl(path string) string {
	ip := config.CoreConf.Sso.SsoIp
	port := config.CoreConf.Sso.SsoPort
	url := "http://" + ip + ":" + utils.GetInterfaceToString(port) + path

	return url
}

func Login(username, password string) string {

	slog.Println(slog.DEBUG, config.CoreConf.Log.LogLevel)

	time.Sleep(1 * time.Second)

	url := getSsoUrl("/tydlpt/auth/login")

	var responseJson LoginRes

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	bodyWriter.WriteField("username", username)
	bodyWriter.WriteField("password", password)

	bodyWriter.Close()
	req, _ := http.NewRequest(http.MethodPost, url, bodyBuf)

	contentType := bodyWriter.FormDataContentType()
	slog.Println(slog.DEBUG, "contentType:", contentType)
	req.Header.Add("Content-Type", contentType)

	cli, addr := GetCli(20 * time.Second)
	resp, err := cli.Do(req)

	if err != nil {
		slog.Println(slog.DEBUG, "登录失败", err)
		return ""
	}

	body, err := ioutil.ReadAll(resp.Body)

	slog.Println(slog.DEBUG, "返回", string(body))
	if err != nil {
		slog.Println(slog.DEBUG, "读取response失败", addr, zap.Error(err))
		return ""
	}
	if err = json.Unmarshal(body, &responseJson); err != nil {
		slog.Println(slog.DEBUG, "json读取失败", addr, zap.Error(err))
		return ""
	}

	if responseJson.ReturnCode != 1 {
		slog.Println(slog.DEBUG, "心跳返回错误", addr, zap.Error(err))
		return ""
	}

	slog.Println(slog.DEBUG, responseJson.Data.InfoMap.Token)

	return responseJson.Data.InfoMap.Token
}
