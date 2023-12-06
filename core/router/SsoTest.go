package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func Login(c *gin.Context) {

	var res LoginRes
	res.ReturnCode = 1
	res.Data.LoginName = c.PostForm("username")
	res.Data.InfoMap.Token = "aaaaaaaaaaaaaaaaaaaaa"
	c.JSON(http.StatusOK, res)
}
