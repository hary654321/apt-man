package probe

import (
	"net/http"
	"zrDispatch/common/log"
	"zrDispatch/core/ginhelp"
	"zrDispatch/core/slog"
	"zrDispatch/core/utils/define"
	"zrDispatch/core/utils/resp"
	"zrDispatch/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateProbeGroup(c *gin.Context) {

	pg := define.ProbeGroupAdd{}

	err := c.ShouldBindJSON(&pg)
	if err != nil {
		log.Error("ShouldBindJSON failed", zap.Error(err))
		resp.JSON(c, resp.ErrBadRequest, nil)
		return
	}

	models.AddProbeGroup(pg)

	resp.JSON(c, resp.Success, nil)
}

func GetProbeGroup(c *gin.Context) {

	query := ginhelp.GetQueryParams(c)

	var q define.Query

	err := c.BindQuery(&q)
	if err != nil {
		slog.Printf(slog.DEBUG, "BindQuery offset failed", err)
	}

	if q.Limit == 0 {
		q.Limit = define.DefaultLimit
	}

	data, count := models.GetProbeGroup(q.Offset, q.Limit, query)

	resp.JSON(c, resp.Success, data, int(count))
}

func DelProbeGroup(c *gin.Context) {

	id := define.GetIDInt{}

	err := c.ShouldBindJSON(&id)
	if err != nil {
		log.Error("ShouldBindJSON failed", zap.Error(err))
		resp.JSON(c, resp.ErrBadRequest, nil)
		return
	}
	data := make(map[string]interface{})

	data["count"] = models.DeleteProbeGroup([]int{id.ID})

	c.JSON(http.StatusOK, gin.H{
		"code": resp.Success,
		"data": data,
	})
}

func EditProbeGroup(c *gin.Context) {

	pg := define.ProbeGroupE{}

	err := c.ShouldBindJSON(&pg)
	if err != nil {
		log.Error("ShouldBindJSON failed", zap.Error(err))
		resp.JSON(c, resp.ErrBadRequest, nil)
		return
	}

	data := make(map[string]interface{})

	data["count"] = models.EditProbeGroupr(pg)

	code := resp.Success
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
	})
}

func GetPgSelect(c *gin.Context) {

	data := models.GetPgSelect()
	code := resp.Success
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
	})
}
