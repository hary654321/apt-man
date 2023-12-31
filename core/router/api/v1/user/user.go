package user

import (
	"context"
	"errors"

	"zrDispatch/common/jwt"
	"zrDispatch/common/utils"
	"zrDispatch/core/client"
	"zrDispatch/core/config"
	"zrDispatch/core/model"
	"zrDispatch/core/slog"
	"zrDispatch/core/utils/define"
	"zrDispatch/core/utils/resp"
	"zrDispatch/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RegistryUser new user
// @Summary registry new user
// @Tags User
// @Produce json
// @Param Registry body define.RegistryUser true "registry user"
// @Success 200 {object} resp.Response
// @Router /api/v1/user/registry [post]
// @Security ApiKeyAuth
func RegistryUser(c *gin.Context) {
	var (
		hashpassword string
	)
	ctx, cancel := context.WithTimeout(context.Background(),
		config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer cancel()
	ruser := define.RegistryUser{}
	err := c.ShouldBindJSON(&ruser)
	if err != nil {
		slog.Println(slog.DEBUG, "ShouldBindJSON failed", zap.Error(err))
		resp.JSON(c, resp.ErrBadRequest, nil)
		return
	}
	// TODO only admin

	hashpassword, err = utils.GenerateHashPass(ruser.Password)
	if err != nil {
		slog.Println(slog.DEBUG, "GenerateHashPass failed", zap.Error(err))
		resp.JSON(c, resp.ErrInternalServer, nil)
		return
	}

	exist, err := model.Check(ctx, model.TBUser, model.Name, ruser.Name)
	if err != nil {
		slog.Println(slog.DEBUG, "IsExist failed", zap.Error(err))
		resp.JSON(c, resp.ErrInternalServer, nil)
		return
	}
	if exist {
		resp.JSON(c, resp.ErrUserNameExist, nil)
		return
	}

	_, err = model.AddUser(ctx, ruser.Name, hashpassword, ruser.Remark, ruser.Role, "")
	if err != nil {
		slog.Println(slog.DEBUG, "AddUser failed", zap.Error(err))
		resp.JSON(c, resp.ErrInternalServer, nil)
		return
	}

	resp.JSON(c, resp.Success, nil)
}

// GetUser Get User Info By Token
// @Summary get user info by token
// @Tags User
// @Produce json
// @Success 200 {object} resp.Response
// @Router /api/v1/user/info [get]
// @Security ApiKeyAuth
func GetUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(),
		config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer cancel()

	uid := c.GetString("uid")
	slog.Println(slog.DEBUG, "uid======", uid)
	exist, err := model.Check(ctx, model.TBUser, model.ID, uid)
	if err != nil {
		slog.Println(slog.DEBUG, "IsExist failed", zap.Error(err))
		resp.JSON(c, resp.ErrInternalServer, nil)
		return
	}
	if !exist {
		resp.JSON(c, resp.ErrUserNotExist, nil)
		return
	}

	user, err := model.GetUserByID(ctx, uid)
	if err != nil {
		slog.Println(slog.DEBUG, "GetUserByID failed", zap.Error(err))
		resp.JSON(c, resp.ErrInternalServer, nil)
		return
	}
	user.Password = ""
	if user.Role == 2 {
		user.Roles = []string{"admin"}
	} else {
		user.Roles = []string{"normal"}
	}
	resp.JSON(c, resp.Success, user)
}

// GetUsers get user info by token
// @Summary  get all users info
// @Tags User
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Produce json
// @Success 200 {object} resp.Response
// @Router /api/v1/user/all [get]
// @Security ApiKeyAuth
func GetUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(),
		config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer cancel()
	var (
		q   define.Query
		err error
	)
	// TODO only admin

	err = c.BindQuery(&q)
	if err != nil {
		slog.Println(slog.DEBUG, "BindQuery offset failed", zap.Error(err))
	}

	if q.Limit == 0 {
		q.Limit = define.DefaultLimit
	}
	users, count, err := model.GetUsers(ctx, nil, q.Offset, q.Limit)
	if err != nil {
		slog.Println(slog.DEBUG, "GetUsers failed", zap.Error(err))
		resp.JSON(c, resp.ErrInternalServer, nil)
		return
	}
	// remove password
	for i, user := range users {
		user.Password = ""
		users[i] = user
	}

	resp.JSON(c, resp.Success, users, count)
}

// ChangeUserInfo change user self config
// @Summary user change self's config info
// @Tags User
// @Description change self config,like email,wechat,dingphone,slack,telegram,password,remark
// @Produce json
// @Param User body define.ChangeUserSelf true "Change Self User Info"
// @Success 200 {object} resp.Response
// @Router /api/v1/user/info [put]
// @Security ApiKeyAuth
func ChangeUserInfo(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(),
		config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer cancel()

	newinfo := define.ChangeUserSelf{}
	err := c.ShouldBindJSON(&newinfo)
	if err != nil {
		slog.Println(slog.DEBUG, "ShouldBindJSON failed", zap.Error(err))
		resp.JSONNew(c, resp.ErrBadRequest, err.Error())
		return
	}
	if len(newinfo.Password) > 0 && len(newinfo.Password) < 8 {
		slog.Println(slog.DEBUG, "password is short 8")
		resp.JSON(c, resp.ErrBadRequest, err.Error())
		return
	}
	uid := c.GetString("uid")
	if uid != newinfo.ID {
		slog.Println(slog.DEBUG, "uid is error", zap.String("uid", uid), zap.String("infoid", newinfo.ID))
		resp.JSON(c, resp.ErrBadRequest, err.Error())
		return
	}
	exist, err := model.Check(ctx, model.TBUser, model.UserName, newinfo.Name, newinfo.ID)
	if err != nil {
		slog.Println(slog.DEBUG, "IsExist failed", zap.Error(err))
		resp.JSON(c, resp.ErrInternalServer, err.Error())
		return
	}
	if exist {
		resp.JSON(c, resp.ErrUserNameExist, err.Error())
		return
	}
	err = model.ChangeUserInfo(ctx,
		uid,
		newinfo.Name,
		newinfo.Password,
		newinfo.Remark)
	if err != nil {
		slog.Println(slog.DEBUG, "ChangeUserInfo failed", zap.Error(err))
		resp.JSON(c, resp.ErrInternalServer, err.Error())
		return
	}

	resp.JSON(c, resp.Success, nil)
}

// AdminChangeUser will change role,forbid,password,Remark
// @Summary admin change user info
// @Tags User
// @Description admin change user's role,forbid,password,remark
// @Produce json
// @Param User body define.AdminChangeUser true "Admin Change User"
// @Success 200 {object} resp.Response
// @Router /api/v1/user/admin [put]
// @Security ApiKeyAuth
func AdminChangeUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(),
		config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer cancel()

	user := define.AdminChangeUser{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		slog.Println(slog.DEBUG, "ShouldBindJSON failed", zap.Error(err))
		resp.JSON(c, resp.ErrBadRequest, err.Error())
		return
	}
	if len(user.Password) > 0 && len(user.Password) < 8 {
		slog.Println(slog.DEBUG, "password is short 8")
		resp.JSON(c, resp.ErrBadRequest, err.Error())
		return
	}
	// TODO only admin
	exist, err := model.Check(ctx, model.TBUser, model.ID, user.ID)
	if err != nil {
		slog.Println(slog.DEBUG, "IsExist failed", zap.Error(err))
		resp.JSON(c, resp.ErrInternalServer, err.Error())
	}
	if !exist {
		resp.JSON(c, resp.ErrUserNotExist, err.Error())
		return
	}
	var role define.Role
	if v, ok := c.Get("role"); ok {
		role = v.(define.Role)
	}
	if role != define.AdminUser {
		resp.JSON(c, resp.ErrAcl, err.Error())
		return
	}

	err = model.AdminChangeUser(ctx, user.ID, user.Role, user.Forbid, user.Password, user.Remark)
	if err != nil {
		slog.Println(slog.DEBUG, "AdminChangeUser failed", zap.Error(err))
		resp.JSON(c, resp.ErrInternalServer, err.Error())
		return
	}

	resp.JSON(c, resp.Success, nil)
}

// AdminDeleteUser will delete user
// @Summary admin delete user
// @Tags User
// @Description admin delet user
// @Produce json
// @Param User body define.AdminChangeUser true "Admin Change User"
// @Success 200 {object} resp.Response
// @Router /api/v1/user/admin [delete]
// @Security ApiKeyAuth
func AdminDeleteUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(),
		config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer cancel()

	user := define.GetID{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		slog.Println(slog.DEBUG, "ShouldBindJSON failed", zap.Error(err))
		resp.JSON(c, resp.ErrBadRequest, nil)
		return
	}

	var role define.Role
	if v, ok := c.Get("role"); ok {
		role = v.(define.Role)
	}
	if role != define.AdminUser {
		resp.JSON(c, resp.ErrAcl, nil)
		return
	}
	// 只能删除普通用户，不能删除admin用户
	// userinfo, err := model.GetUserByID(ctx, user.ID)
	// if err != nil {
	// 	slog.Println(slog.DEBUG,"GetUserByID failed", zap.Error(err))
	// 	resp.JSON(c, resp.ErrInternalServer, nil)
	// 	return
	// }
	// if userinfo.Role == define.AdminUser {
	// 	resp.JSON(c, resp.ErrAcl, nil)
	// 	return
	// }
	// TODO only admin
	exist, err := model.Check(ctx, model.TBUser, model.ID, user.ID)
	if err != nil {
		slog.Println(slog.DEBUG, "IsExist failed", zap.Error(err))
		resp.JSON(c, resp.ErrInternalServer, nil)
	}
	if !exist {
		resp.JSON(c, resp.ErrUserNotExist, nil)
		return
	}
	// 检查用户是否创建资源
	ok1, err := model.Check(ctx, model.TBTask, model.CreateByID, user.ID)
	if err != nil {
		slog.Println(slog.DEBUG, "Check failed", zap.Error(err))
		resp.JSON(c, resp.ErrInternalServer, nil)
		return
	}
	ok2, err := model.Check(ctx, model.TBHostgroup, model.CreateByID, user.ID)
	if err != nil {
		slog.Println(slog.DEBUG, "Check failed", zap.Error(err))
		resp.JSON(c, resp.ErrInternalServer, nil)
		return
	}
	if ok1 || ok2 {
		resp.JSON(c, resp.ErrDelUserUseByOther, nil)
		return
	}
	err = model.DeleteUser(ctx, user.ID)
	if err != nil {
		slog.Println(slog.DEBUG, "DeleteUser failed", zap.Error(err))
		resp.JSON(c, resp.ErrInternalServer, nil)
		return
	}
	resp.JSON(c, resp.Success, nil)
}

// LoginUser login user
// @Summary login user
// @Tags User
// @Produce json
// @Success 200 {object} resp.Response
// @Router /api/v1/user/login [post]
// @Security BasicAuth
func LoginUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(),
		config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer cancel()
	username, password, ok := c.Request.BasicAuth()
	if !ok {
		resp.JSON(c, resp.ErrBadRequest, nil)
		return
	}
	token, err := model.LoginUser(ctx, username, password)
	if err != nil {
		slog.Println(slog.DEBUG, "model.LoginUser failed", zap.Error(err))
	}
	switch err := errors.Unwrap(err); err.(type) {
	case nil:
		resp.JSON(c, resp.Success, token)
	case define.ErrUserPass:
		resp.JSON(c, resp.ErrUserPassword, nil)
	case define.ErrForbid:
		resp.JSON(c, resp.ErrUserForbid, nil)
	default:
		resp.JSON(c, resp.ErrInternalServer, nil)
	}
}

func Login(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(),
		config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer cancel()

	ruser := define.Login{}
	err := c.ShouldBindJSON(&ruser)
	if err != nil {
		slog.Println(slog.DEBUG, "ShouldBindJSON failed", zap.Error(err))
		resp.JSON(c, resp.ErrBadRequest, nil)
		return
	}

	token, err := model.LoginUser(ctx, ruser.Name, ruser.Password)
	if err != nil {
		slog.Println(slog.DEBUG, "model.LoginUser failed", zap.Error(err))
	}
	switch err := errors.Unwrap(err); err.(type) {
	case nil:
		resp.JSON(c, resp.Success, token)
	case define.ErrUserPass:
		resp.JSON(c, resp.ErrUserPassword, nil)
	case define.ErrForbid:
		resp.JSON(c, resp.ErrUserForbid, nil)
	default:
		resp.JSON(c, resp.ErrInternalServer, nil)
	}
}
func SsoLogin(c *gin.Context) {

	id := c.Query("userId")
	token := c.Query("token")
	// resourceId := c.Query("resourceId")

	user, err := models.GetUserInfoBYid(id)
	if err != nil {
		resp.JSONNew(c, resp.AddFail, "不存在此用户")
		return
	}

	if !client.CheckToken(token) {
		resp.JSONNew(c, resp.AddFail, "无效token")
		return
	}

	tokennew, err := jwt.GenerateToken(id, user.Name)

	if err != nil {
		slog.Println(slog.DEBUG, "model.LoginUser failed", zap.Error(err))
	}
	switch err := errors.Unwrap(err); err.(type) {
	case nil:
		resp.JSON(c, resp.Success, tokennew)
	case define.ErrUserPass:
		resp.JSON(c, resp.ErrUserPassword, nil)
	case define.ErrForbid:
		resp.JSON(c, resp.ErrUserForbid, nil)
	default:
		resp.JSON(c, resp.ErrInternalServer, nil)
	}
}

// LogoutUser logout user
// @Summary logout user
// @Tags User
// @Produce json
// @Success 200 {object} resp.Response
// @Router /api/v1/user/logout [post]
// @Security BasicAuth
func LogoutUser(c *gin.Context) {
	resp.JSON(c, resp.Success, nil)
}

// GetSelect return name,id
// @Summary return name,id
// @Produce json
// @Success 200 {object} resp.Response
// @Router /api/v1/user/select [post]
// @Security BasicAuth
func GetSelect(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(),
		config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer cancel()
	data, err := model.GetNameID(ctx, model.TBUser)
	if err != nil {
		slog.Println(slog.DEBUG, "model.GetNameID failed", zap.Error(err))
		resp.JSON(c, resp.ErrInternalServer, nil)
		return
	}
	resp.JSON(c, resp.Success, data)
}

// GetAlarmStatus return enable alarm notify
func GetAlarmStatus(c *gin.Context) {
	type NotifyStatus struct {
		Email    bool `json:"email"`
		DingDing bool `json:"dingphone"`
		Slack    bool `json:"slack"`
		Telegram bool `json:"telegram"`
		WeChat   bool `json:"wechat"`
		WebHook  bool `json:"wehook"`
	}
	notifycfg := config.CoreConf.Notify
	notifystatus := NotifyStatus{
		Email:    notifycfg.Email.Enable,
		DingDing: notifycfg.DingDing.Enable,
		Slack:    notifycfg.Slack.Enable,
		Telegram: notifycfg.Telegram.Enable,
		WeChat:   notifycfg.WeChat.Enable,
		WebHook:  notifycfg.WebHook.Enable,
	}
	resp.JSON(c, resp.Success, notifystatus)
}

// GetOperateLog get user operate log
func GetOperateLog(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(),
		config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer cancel()
	type queryparams struct {
		define.Query
		UserName string `form:"username"`
		Method   string `form:"method"`
		Module   string `form:"module"`
	}

	q := queryparams{}
	err := c.ShouldBindQuery(&q)
	if err != nil {
		resp.JSON(c, resp.ErrBadRequest, nil)
		return
	}
	if q.Limit == 0 {
		q.Limit = define.DefaultLimit
	}

	// uid, method, module, limit, offset
	oplogs, count, err := model.GetOperate(ctx, "", q.UserName, q.Method, q.Module, q.Limit, q.Offset)
	if err != nil {
		slog.Println(slog.DEBUG, "model.GetOperate filed", zap.Error(err))
		resp.JSON(c, resp.ErrInternalServer, nil)
		return
	}
	resp.JSON(c, resp.Success, oplogs, count)
}
