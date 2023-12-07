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

	res.Data = append(res.Data, client.UserInfo{Name: "管理员", LoginName: "admin"})

	c.JSON(http.StatusOK, res)
}
