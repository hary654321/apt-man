package router

import (
	"net/http"
	"zrDispatch/core/client"
	"zrDispatch/core/slog"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var res client.LoginRes
	res.ReturnCode = 1
	res.Data.LoginName = c.PostForm("username")
	res.Data.Name = c.PostForm("password")
	res.Data.InfoMap.Token = "aaaaaaaaaaaaaaaaaaaaa"
	c.JSON(http.StatusOK, res)
}

func UserLIst(c *gin.Context) {

	params := &client.UserListReq{}
	if err := c.BindJSON(params); err != nil {
		slog.Println(slog.DEBUG, err)
	}

	slog.Println(slog.DEBUG, params)

	var res client.UserListRes
	res.ReturnCode = 1
	res.Data = []client.UserInfo{}

	res.Data = append(res.Data, client.UserInfo{Name: "管理员", LoginName: "admin", ID: "1111111111"})

	c.JSON(http.StatusOK, res)
}

func CheckLogin(c *gin.Context) {

	c.PostForm("token")

	slog.Println(slog.DEBUG, "token")
	var res client.CheckRes
	res.ReturnCode = 200
	res.ErrorMsg = ""
	res.Data = true

	c.JSON(http.StatusOK, res)
}
