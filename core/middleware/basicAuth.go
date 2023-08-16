package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		token := c.GetHeader("Authorization")
		//slog.Printf(slog.INFO, "Authorization %s  %s", token, global.ServerSetting.BasicAuth)
		if token != "" {
			code = 401
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "验签不通过",
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
