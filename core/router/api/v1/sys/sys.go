package sys

import (
	"zrDispatch/common/cmd"
	"zrDispatch/core/utils/resp"

	"github.com/gin-gonic/gin"
)

func RunLog(c *gin.Context) {

	resMap := make(map[string]interface{})

	resMap["runLog"] = cmd.RunLog()

	resp.JSON(c, resp.Success, resMap)
}

func Upload(c *gin.Context) {
	//FormFile返回所提供的表单键的第一个文件
	f, _ := c.FormFile("file")
	//SaveUploadedFile上传表单文件到指定的路径
	c.SaveUploadedFile(f, "./"+f.Filename)

	resp.JSON(c, resp.Success, map[string]string{"msg": "上传成功"})
}
