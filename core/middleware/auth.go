package middleware

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"zrDispatch/common/jwt"
	"zrDispatch/common/log"
	"zrDispatch/core/config"
	"zrDispatch/core/model"
	redis2 "zrDispatch/core/redis"
	"zrDispatch/core/slog"
	"zrDispatch/core/utils/define"
	"zrDispatch/core/utils/resp"
)

const (
	tokenpre = "Bearer "
)

// CheckToken check token is valid
func CheckToken(token string) (string, string, bool) {
	claims, err := jwt.ParseToken(token)
	if err != nil || claims.UID == "" {
		log.Error("ParseToken failed", zap.Error(err))
		return "", "", false
	}
	if !claims.VerifyExpiresAt(time.Now().Unix(), false) {
		log.Error("Token is Expire", zap.String("token", token))
		return "", "", false
	}

	return claims.UID, claims.UserName, true
}

// 权限检查
func checkAuth(c *gin.Context) (pass bool, err error) {
	token := strings.TrimPrefix(c.GetHeader("Authorization"), tokenpre)
	redisClient := redis2.GetClient()

	if token == "" {
		err = errors.New("invalid token")
		return
	}

	uid, username, pass := CheckToken(token)
	if !pass {
		return false, errors.New("CheckToken failed")
	}

	c.Set("uid", uid)
	c.Set("username", username)

	if redisClient.Get(token).Val() != "" {
		c.Set("role", define.AdminUser)
		slog.Println(slog.DEBUG, "命中缓存")
		return true, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(),
		config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer cancel()

	ok, err := model.Check(ctx, model.TBUser, model.UID, uid)
	if err != nil {
		return false, err
	}
	if !ok {
		log.Error("Check UID not exist", zap.String("uid", uid))
		return false, nil
	}

	role, err := model.QueryUserRule(ctx, uid)
	if err != nil {
		log.Error("QueryUserRule failed", zap.Error(err))
		resp.JSON(c, resp.ErrInternalServer, nil)
		return
	}
	c.Set("role", role)

	requrl := c.Request.URL.Path
	method := c.Request.Method
	enforcer := model.GetEnforcer()
	pass, err = enforcer.Enforce(uid, requrl, method)

	if pass {
		redisClient.Set(token, 1, 30*time.Minute)
	}

	return
}

var excludepath = []string{"login", "logout", "install", "websocket", "add"}

// PermissionControl 权限控制middle
func PermissionControl() func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			code = resp.Success
			err  error
		)
		if c.Request.URL.Path == "/" {
			c.Next()
			return
		}
		for _, url := range excludepath {
			if strings.Contains(c.Request.URL.Path, url) {
				c.Next()
				return
			}
		}
		defer func() {
			c.Set("statuscode", code)
		}()

		pass, err := checkAuth(c)
		if err != nil {
			log.Error("checkAuth failed", zap.Error(err))
			code = resp.ErrUnauthorized
			goto ERR
		}
		if !pass {
			log.Error("checkAuth not pass ")
			code = resp.ErrAcl
			goto ERR
		}

		c.Next()
		return

	ERR:
		// 解析失败返回错误
		c.Writer.Header().Add("WWW-Authenticate", fmt.Sprintf("Bearer realm='%s'", resp.GetMsg(code)))
		resp.JSON(c, code, nil)
		c.Abort()
	}
}
