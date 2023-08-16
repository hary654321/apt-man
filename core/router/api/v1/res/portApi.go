package res

import (
	"net/http"
	"zrDispatch/common/utils"
	"zrDispatch/core/ginhelp"
	"zrDispatch/core/slog"
	"zrDispatch/core/utils/define"
	"zrDispatch/core/utils/resp"
	"zrDispatch/models"

	"github.com/gin-gonic/gin"
)

func GetPoertRes(c *gin.Context) {

	query := ginhelp.GetQueryParams(c)

	var q define.Query

	err := c.BindQuery(&q)
	if err != nil {
		slog.Printf(slog.DEBUG, "BindQuery offset failed", err)
	}

	if q.Limit == 0 {
		q.Limit = define.DefaultLimit
	}

	data, count := models.GetPortRes(q.Offset, q.Limit, query)

	resp.JSON(c, resp.Success, data, int(count))
}

func GetOsSelect(c *gin.Context) {

	data := models.GetOsSelect()
	resp.JSON(c, resp.Success, data)
}

func UpdatePortRemark(c *gin.Context) {

	pg := define.PortResEdit{}

	err := c.ShouldBindJSON(&pg)
	if err != nil {
		slog.Println(slog.DEBUG, "ShouldBindJSON failed", err)
		resp.JSON(c, resp.ErrBadRequest, nil)
		return
	}

	data := make(map[string]interface{})

	pg.Utime = utils.GetTimeStr()
	data["count"] = models.EditPortRes(pg)

	code := resp.Success
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
	})
}
