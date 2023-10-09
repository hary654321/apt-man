package plug

import (
	"net/http"
	"strings"

	"zrDispatch/common/cmd"
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

func CreatePlug(c *gin.Context) {

	f, _ := c.FormFile("file")
	//SaveUploadedFile上传表单文件到指定的路径
	c.SaveUploadedFile(f, "/app/"+f.Filename)

	slog.Println(slog.DEBUG, f.Header)

	pi := define.PlugInfoAdd{
		Name:     c.PostForm("name"),
		Desc:     c.PostForm("desc"),
		Cmd:      c.PostForm("cmd"),
		FileName: f.Filename,
		Status:   define.PLUG_JYZ,
	}

	res := models.GetPlugInfoByName(pi.Name, "")

	if res.Name != "" {
		resp.JSON(c, resp.PnameExits, nil)
		return
	}

	err1 := models.AddPlugInfo(pi)

	if err1 != nil {
		resp.JSON(c, resp.AddFail, err1.Error())
		return
	}

	go checkPlug(pi.Name, f.Filename, pi.Cmd)

	resp.JSON(c, resp.Success, nil)
}

// 必须要带参数才可以
func checkPlug(pname, filename, cmdstr string) {
	cmd.Exec("chmod +x " + "/app/" + filename)

	cmdstr = "/app/" + cmdstr

	cmdstr = strings.Replace(cmdstr, "{ip}", "127.0.0.1", -1)

	cmdstr = strings.Replace(cmdstr, "{port}", "80", -1)

	cmdstr = strings.Replace(cmdstr, "{res}", "test.res", -1)

	_, err := cmd.Exec("/app/" + filename)

	if err != nil {
		slog.Println(slog.DEBUG, err)
		models.ChangePlugState(pname, int(define.PLUG_FAIL))
	} else {
		models.ChangePlugState(pname, int(define.PLUG_SUCC))
	}

}

func GetPlug(c *gin.Context) {

	query := ginhelp.GetQueryParams(c)

	var q define.Query

	err := c.BindQuery(&q)
	if err != nil {
		slog.Printf(slog.DEBUG, "BindQuery offset failed", err)
	}

	if q.Limit == 0 {
		q.Limit = define.DefaultLimit
	}

	data, count := models.GetPlugInfo(q.Offset, q.Limit, query)

	resp.JSON(c, resp.Success, data, int(count))
}

func DelPlug(c *gin.Context) {

	id := define.GetIDInt{}

	err := c.ShouldBindJSON(&id)
	if err != nil {
		log.Error("ShouldBindJSON failed", zap.Error(err))
		resp.JSON(c, resp.ErrBadRequest, nil)
		return
	}
	data := make(map[string]interface{})
	data["count"] = models.DeletePlugInfo([]int{id.ID})

	c.JSON(http.StatusOK, gin.H{
		"code": resp.Success,
		"data": data,
	})
}

func EditPlug(c *gin.Context) {

	pi := define.PlugInfoE{
		Name: c.PostForm("name"),
		Desc: c.PostForm("desc"),
		Cmd:  c.PostForm("cmd"),
		ID:   c.PostForm("id"),
	}

	res := models.GetPlugInfoByName(pi.Name, pi.ID)

	if res.Name != "" {
		resp.JSON(c, resp.PnameExits, nil)
		return
	}

	f, _ := c.FormFile("file")

	if f != nil {
		pi.FileName = f.Filename
		//SaveUploadedFile上传表单文件到指定的路径
		c.SaveUploadedFile(f, "/app/"+f.Filename)

	}

	data := make(map[string]interface{})

	serr := models.EditPlugInfo(pi)

	if serr != nil {
		resp.JSON(c, resp.AddFail, serr.Error())
		return
	}
	if f != nil {
		go checkPlug(pi.Name, f.Filename)
	}
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

	res := models.BatchAddPlugInfo(data)

	if res == 0 {
		resp.JSON(c, resp.AddFail, nil)
		return
	}

	resp.JSON(c, resp.Success, map[string]string{"msg": "导入成功"})
}

func GetPiSelect(c *gin.Context) {

	data := models.GetPlugSelect()
	code := resp.Success
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
