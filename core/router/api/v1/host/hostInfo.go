package host

import (
	"context"
	"github.com/gin-gonic/gin"
	"zrDispatch/common/log"
	"zrDispatch/common/utils"
	"zrDispatch/core/client"
	"zrDispatch/core/config"
	"zrDispatch/core/model"
	"zrDispatch/core/utils/resp"
)

// ChangeHostState stop host worker
// @Summary stop host worker
// @Tags Host
// @Description stop host worker
// @Param StopHost body define.GetID true "ID"
// @Produce json
// @Success 200 {object} resp.Response
// @Router /api/v1/host/stop [put]
// @Security ApiKeyAuth
func GetHostInfo(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(),
		config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer cancel()

	hostId := c.Query("id")

	println(hostId)
	if utils.CheckID(hostId) != nil {
		log.Error("CheckID failed")
		resp.JSON(c, resp.ErrBadRequest, nil)
		return
	}
	host, _ := model.GetHostByID(ctx, hostId)

	hostInfo := client.HostInfo(host)

	resp.JSON(c, resp.Success, hostInfo)

}
