package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
	"zrDispatch/common/utils"
	"zrDispatch/core/config"
	"zrDispatch/core/model"
	"zrDispatch/core/slog"
	"zrDispatch/core/utils/define"

	"github.com/bytedance/sonic"
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
	url := config.CoreConf.Sso.Secm + "://" + ip + ":" + utils.GetInterfaceToString(port) + path

	return url
}

func Login(username, password string) string {

	url := getSsoUrl("/tydlpt/auth/login")

	var responseJson LoginRes

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	bodyWriter.WriteField("username", username)
	bodyWriter.WriteField("password", password)

	bodyWriter.Close()
	req, _ := http.NewRequest(http.MethodPost, url, bodyBuf)

	contentType := bodyWriter.FormDataContentType()
	// slog.Println(slog.DEBUG, "contentType:", contentType)
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

type UserListRes struct {
	ReturnCode int         `json:"returnCode"`
	Data       []UserInfo  `json:"data"`
	ErrorMsg   interface{} `json:"errorMsg"`
}

type UserInfo struct {
	BeginLastUpdateTime string `json:"beginLastUpdateTime"`
	Birthday            string `json:"birthday"`
	EndLastUpdateTime   string `json:"endLastUpdateTime"`
	SystemGlobalID      string `json:"systemGlobalId"`
	Roles               []struct {
		UpdaterID   int    `json:"updaterId"`
		AppID       int    `json:"appId"`
		DisableFlag bool   `json:"disableFlag"`
		GlobalID    string `json:"globalId"`
		Name        string `json:"name"`
		Description string `json:"description"`
		UpdateTime  int64  `json:"updateTime"`
		ID          int    `json:"id"`
		Value       string `json:"value"`
	} `json:"roles"`
	Sex             string        `json:"sex"`
	DeptID          string        `json:"deptId"`
	ExtensionFields []interface{} `json:"extensionFields"`
	PhoneNumber     string        `json:"phoneNumber"`
	DisableFlag     bool          `json:"disableFlag"`
	LoginName       string        `json:"loginName"`
	Name            string        `json:"name"`
	ID              string        `json:"id"`
	UserType        string        `json:"userType"`
	LastUpdateTime  string        `json:"lastUpdateTime"`
}

type UserListReq struct {
	Token string `json:"token"`
}

func GetAllUserList() (res UserListRes) {
	time.Sleep(1 * time.Second)

	url := getSsoUrl("/tydlpt/auth/listAllUserInfo")

	var reqD UserListReq
	reqD.Token = Login("a", "b")

	jsonData, _ := sonic.Marshal(&reqD)
	slog.Println(slog.DEBUG, "发送", string(jsonData))

	body, err := Send(url, string(jsonData))

	slog.Println(slog.DEBUG, "返回", string(body))
	if err != nil {
		slog.Println(slog.DEBUG, "读取response失败", url, zap.Error(err))
		return
	}
	if err = json.Unmarshal(body, &res); err != nil {
		slog.Println(slog.DEBUG, "json读取失败", url, zap.Error(err))
		return
	}

	if res.ReturnCode != 1 {
		slog.Println(slog.DEBUG, "心跳返回错误", url, zap.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(),
		config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer cancel()
	for _, v := range res.Data {

		slog.Println(slog.DEBUG, "sso导入", v.LoginName)
		hashpassword, _ := utils.GenerateHashPass("zrtx@2023")
		_, err = model.AddUser(ctx, v.LoginName, hashpassword, "sso导入", define.AdminUser, v.ID)
	}

	return
}

type CheckRes struct {
	ReturnCode int    `json:"returnCode"`
	Data       bool   `json:"data"`
	ErrorMsg   string `json:"errorMsg"`
}

func CheckToken(token string) bool {
	slog.Println(slog.DEBUG, config.CoreConf.Log.LogLevel)

	time.Sleep(1 * time.Second)

	url := getSsoUrl("/tydlpt/auth/authToken")

	var responseJson CheckRes

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	bodyWriter.WriteField("token", token)

	bodyWriter.Close()
	req, _ := http.NewRequest(http.MethodPost, url, bodyBuf)

	contentType := bodyWriter.FormDataContentType()
	// slog.Println(slog.DEBUG, "contentType:", contentType)
	req.Header.Add("Content-Type", contentType)

	cli, addr := GetCli(20 * time.Second)
	resp, err := cli.Do(req)

	if err != nil {
		slog.Println(slog.DEBUG, "登录失败", err)
		return false
	}

	body, err := ioutil.ReadAll(resp.Body)

	slog.Println(slog.DEBUG, "返回", string(body))
	if err != nil {
		slog.Println(slog.DEBUG, "读取response失败", addr, zap.Error(err))
		return false
	}
	if err = json.Unmarshal(body, &responseJson); err != nil {
		slog.Println(slog.DEBUG, "json读取失败", addr, zap.Error(err))
		return false
	}

	if responseJson.ReturnCode != 200 {
		slog.Println(slog.DEBUG, "check 失败", addr, zap.Error(err))
		return false
	}

	return true
}
