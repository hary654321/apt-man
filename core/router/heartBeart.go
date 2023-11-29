package router

import (
	"net/http"
	"zrDispatch/common/utils"

	"github.com/gin-gonic/gin"
)

type HeartBeatS struct {
	time         string
	version      string
	runningTasks string
}

func HeartBeat(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"time":        utils.GetTime(),
		"name":        "apt-server",
		"status":      1,
		"description": "",
	})
}
