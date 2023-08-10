package res

import (
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
