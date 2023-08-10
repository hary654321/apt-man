package probe

import (
	"net/http"
	"zrDispatch/common/log"
	"zrDispatch/common/utils"
	"zrDispatch/core/ginhelp"
	"zrDispatch/core/slog"
	"zrDispatch/core/utils/define"
	"zrDispatch/core/utils/resp"
	"zrDispatch/e"
	"zrDispatch/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateProbe(c *gin.Context) {

	pi := define.ProbeInfoAdd{}

	err := c.ShouldBindJSON(&pi)
	if err != nil {
		log.Error("ShouldBindJSON failed", zap.Error(err))
		resp.JSON(c, resp.ErrBadRequest, nil)
		return
	}

	res := models.GetProbeInfoByName(pi.Name)

	utils.WriteJsonLog(res)
	if res.Name != "" {
		resp.JSON(c, resp.PnameExits, nil)
		return
	}

	err = models.AddProbeInfo(pi)

	if err != nil {
		resp.JSON(c, resp.AddFail, err.Error())
		return
	}

	resp.JSON(c, resp.Success, nil)
}

func GetProbe(c *gin.Context) {

	query := ginhelp.GetQueryParams(c)

	var q define.Query

	err := c.BindQuery(&q)
	if err != nil {
		slog.Printf(slog.DEBUG, "BindQuery offset failed", err)
	}

	if q.Limit == 0 {
		q.Limit = define.DefaultLimit
	}

	data, count := models.GetProbeInfo(q.Offset, q.Limit, query)

	resp.JSON(c, resp.Success, data, int(count))
}

func DelProbe(c *gin.Context) {

	id := define.GetIDInt{}

	err := c.ShouldBindJSON(&id)
	if err != nil {
		log.Error("ShouldBindJSON failed", zap.Error(err))
		resp.JSON(c, resp.ErrBadRequest, nil)
		return
	}
	data := make(map[string]interface{})
	data["count"] = models.DeleteProbeInfo([]int{id.ID})

	c.JSON(http.StatusOK, gin.H{
		"code": resp.Success,
		"data": data,
	})
}

func EditProbe(c *gin.Context) {

	pg := define.ProbeInfoE{}

	err := c.ShouldBindJSON(&pg)
	if err != nil {
		log.Error("ShouldBindJSON failed", zap.Error(err))
		resp.JSON(c, resp.ErrBadRequest, nil)
		return
	}

	data := make(map[string]interface{})

	data["count"] = models.EditProbeInfo(pg)

	code := resp.Success
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
	})
}

func Import(c *gin.Context) {
	//FormFile返回所提供的表单键的第一个文件
	f, _ := c.FormFile("file")
	//SaveUploadedFile上传表单文件到指定的路径
	c.SaveUploadedFile(f, "./"+f.Filename)

	data, _ := utils.GetcsvDataPro(f.Filename)

	res := models.BatchAddProbeInfo(data)

	if res == 0 {
		resp.JSON(c, resp.AddFail, nil)
		return
	}

	resp.JSON(c, resp.Success, map[string]string{"msg": "导入成功"})
}

func GetPiSelect(c *gin.Context) {

	data := models.GetProbeSelect()
	code := resp.Success
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
