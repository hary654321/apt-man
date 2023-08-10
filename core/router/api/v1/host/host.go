package host

import (
	"context"
	"time"
	"zrDispatch/common/log"
	"zrDispatch/common/utils"
	"zrDispatch/core/config"
	"zrDispatch/core/model"
	"zrDispatch/core/service"
	"zrDispatch/core/slog"
	sshclient "zrDispatch/core/ssh"
	"zrDispatch/core/utils/define"
	"zrDispatch/core/utils/resp"
	"zrDispatch/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetHost return all registry gost
// @Summary get all hosts
// @Tags Host
// @Description get all registry host
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Produce json
// @Success 200 {object} resp.Response
// @Router /api/v1/host [get]
// @Security ApiKeyAuth
func GetHost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(),
		config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer cancel()
	var (
		q   define.Query
		err error
	)

	err = c.BindQuery(&q)
	if err != nil {
		log.Error("BindQuery offset failed", zap.Error(err))
	}

	if q.Limit == 0 {
		q.Limit = define.DefaultLimit
	}

	hosts, count, err := model.GetHosts(ctx, q.Offset, q.Limit)

	if err != nil {
		log.Error("GetHost failed", zap.Error(err))
		resp.JSON(c, resp.ErrInternalServer, nil)
		return
	}

	resp.JSON(c, resp.Success, hosts, count)
}

func CreateHost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(),
		config.CoreConf.Server.DB.MaxQueryTime.Duration)
	//config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer cancel()

	hostInfo := model.HostInfo{}
	err := c.ShouldBindJSON(&hostInfo)
	if err != nil {
		log.Error("ShouldBindJSON failed", zap.Error(err))
		resp.JSON(c, resp.ErrBadRequest, nil)
		return
	}

	id := utils.GetID()

	id, err = model.CreateHost(ctx, hostInfo)

	if err != nil {
		log.Error("CreateHost failed", zap.Error(err))
		resp.JSON(c, resp.ErrHostExist, nil)
		return
	}
	log.Info(id)

	resp.JSON(c, resp.Success, nil)
}

func ChangeHost(c *gin.Context) {

	hostInfo := define.HostGorm{}
	err := c.ShouldBindJSON(&hostInfo)
	if err != nil {
		log.Error("ShouldBindJSON failed", zap.Error(err))
		resp.JSON(c, resp.ErrBadRequest, nil)
		return
	}

	err = models.ChangeHost(hostInfo)

	if err != nil {
		log.Error("changehost failed", zap.Error(err))
		resp.JSON(c, resp.ErrHostExist, nil)
		return
	}

	resp.JSON(c, resp.Success, nil)
}

// ChangeHostState stop host worker
// @Summary stop host worker
// @Tags Host
// @Description stop host worker
// @Param StopHost body define.GetID true "ID"
// @Produce json
// @Success 200 {object} resp.Response
// @Router /api/v1/host/stop [put]
// @Security ApiKeyAuth
func ChangeHostState(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(),
		config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer cancel()
	gethost := define.GetID{}
	err := c.ShouldBindJSON(&gethost)
	if err != nil {
		log.Error("c.ShouldBindJSON", zap.Error(err))
		resp.JSON(c, resp.ErrBadRequest, nil)
		return
	}
	if utils.CheckID(gethost.ID) != nil {
		log.Error("CheckID failed")
		resp.JSON(c, resp.ErrBadRequest, nil)
		return
	}
	host, err := model.GetHostByID(ctx, gethost.ID)
	switch err.(type) {
	case nil:
		goto Next
	case define.ErrNotExist:
		resp.JSON(c, resp.ErrHostNotExist, nil)
		return
	default:

		resp.JSON(c, resp.ErrInternalServer, nil)
		return
	}
Next:
	err = model.StopHost(ctx, gethost.ID, !host.Stop)
	if err != nil {
		log.Error("model.StopHost", zap.Error(err))
		resp.JSON(c, resp.ErrInternalServer, nil)
		return
	}
	resp.JSON(c, resp.Success, nil)
}

// DeleteHost delete host
// @Summary delete host
// @Tags Host
// @Description delete host
// @Param StopHost body define.GetID true "ID"
// @Produce json
// @Success 200 {object} resp.Response
// @Router /api/v1/host [delete]
// @Security ApiKeyAuth
func DeleteHost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(),
		config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer cancel()
	gethost := define.GetID{}
	err := c.ShouldBindJSON(&gethost)
	if err != nil {
		resp.JSON(c, resp.ErrBadRequest, nil)
		return
	}
	if utils.CheckID(gethost.ID) != nil {
		resp.JSON(c, resp.ErrBadRequest, nil)
		return
	}

	hostgroups, _, err := model.GetHostGroups(ctx, 0, 0)
	if err != nil {
		resp.JSON(c, resp.ErrInternalServer, nil)
		return
	}
	for _, hostgroup := range hostgroups {
		for _, hid := range hostgroup.HostsID {
			if gethost.ID == hid {
				resp.JSON(c, resp.ErrDelHostUseByOtherHG, nil)
				return
			}
		}
	}

	err = model.DeleteHost(ctx, gethost.ID)
	if err != nil {
		log.Error("model.DeleteHost failed", zap.Error(err))
		resp.JSON(c, resp.ErrInternalServer, nil)
		return
	}
	resp.JSON(c, resp.Success, nil)
}

// GetSelect name,id
// @Summary Get Task Select
// @Tags Host
// @Produce json
// @Success 200 {object} resp.Response
// @Router /api/v1/host/select [get]
// @Security ApiKeyAuth
func GetSelect(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(),
		config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer cancel()
	data, err := model.GetNameID(ctx, model.TBHost)
	if err != nil {
		log.Error("model.GetNameID failed", zap.Error(err))
		resp.JSON(c, resp.ErrInternalServer, nil)
		return
	}
	resp.JSON(c, resp.Success, data)
}

func StartService(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(),
		config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer cancel()

	id := c.Query("id")
	slog.Println(slog.DEBUG, "id", id)
	host, _ := model.GetHostByID(ctx, id)

	s, res := sshclient.Start(host, true)

	if res == true {
		resp.JSON(c, resp.Success, s)
	} else {
		resp.JSON(c, resp.ErrInternalServer, s)
	}

}

// 服务日志
func ServiceLog(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(),
		config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer cancel()

	id := c.Query("id")
	slog.Println(slog.DEBUG, "id", id)
	println(id)
	host, _ := model.GetHostByID(ctx, id)

	log, res := sshclient.ServiceLog(host)

	serviceCmd, _ := sshclient.ServiceCmd(host)

	resMap := make(map[string]interface{})

	resMap["log"] = log
	resMap["serviceCmd"] = serviceCmd

	if res == true {
		resp.JSON(c, resp.Success, resMap)
	} else {
		resp.JSON(c, resp.ErrInternalServer, resMap)
	}

}

func Bash(c *gin.Context) {

	gid := c.Query("gid")
	bash := c.Query("bash")

	ctx, cancel := context.WithTimeout(context.Background(),
		config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer cancel()

	hosts, _ := model.GetHostsByHGID(ctx, gid)
	// slog.Println(slog.DEBUG, gid, hosts)
	count := len(hosts)

	if count < 1 {
		return
	}

	for _, hostInfo := range hosts {

		go service.Bash(bash, hostInfo)

		time.Sleep(1 * time.Second)

	}

	resp.JSON(c, resp.Success, hosts, count)
}

func CleanLog(c *gin.Context) {

	gid := c.Query("gid")

	ctx, cancel := context.WithTimeout(context.Background(),
		config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer cancel()

	hosts, _ := model.GetHostsByHGID(ctx, gid)
	// slog.Println(slog.DEBUG, gid, hosts)
	count := len(hosts)

	if count < 1 {
		return
	}

	for _, hostInfo := range hosts {
		go sshclient.CleanLog(hostInfo)
		time.Sleep(1 * time.Second)
	}

	resp.JSON(c, resp.Success, hosts, count)
}

func BinDeploy(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(),
		config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer cancel()

	id := c.Query("id")
	slog.Println(slog.DEBUG, "id", id)
	hostInfo, err := model.GetHostByID(ctx, id)

	if err != nil {
		resp.JSON(c, resp.InstallFail, err.Error())
		return
	}

	err = BinInstall(hostInfo)
	if err != nil {
		resp.JSON(c, resp.InstallFail, err.Error())
		return
	}

	resp.JSON(c, resp.Success, hostInfo)
}

func BinInstall(hostInfo *define.Host) (err error) {
	client, err := sshclient.DialWithPasswd(hostInfo.Ip+":"+utils.GetInterfaceToString(hostInfo.SshPort), hostInfo.SshUser, hostInfo.SshPwd)
	if err != nil {
		slog.Println(slog.DEBUG, "sshclient.DialWithPasswd failed", zap.Error(err))
		return
	}
	defer client.Close()

	client.Cmd("mkdir -p /zrtx/apt").Output()

	client.Upload("./client.tar", "/zrtx/apt/client.tar")

	client.Cmd("tar -xvf /zrtx/apt/client.tar -C /zrtx/apt").Output()

	client.Cmd("chmod +x /zrtx/apt/clientInstall.sh").Output()
	go client.Cmd("/zrtx/apt/clientInstall.sh").Output()

	models.UpdateHostStatus(hostInfo.ID, model.RUNNING)

	return nil
}

func Restart(c *gin.Context) {

	//utils.ZipFile("../scanning-client", "./update.zip", "2023-02-13 15:04:05")

	ctx, cancel := context.WithTimeout(context.Background(),
		config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer cancel()
	hosts, count, err := model.GetHostsWithStatus(ctx, "=", model.RUNNING)

	if err != nil {
		log.Error("BindQuery offset failed", zap.Error(err))
	}

	if count < 1 {
		return
	}

	for _, hostInfo := range hosts {

		go sshclient.Restart(hostInfo)

	}
	resp.JSON(c, resp.Success, hosts, count)
}
