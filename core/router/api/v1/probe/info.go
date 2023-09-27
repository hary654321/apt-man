package probe

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
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
		resp.JSON(c, resp.ProbeInfoAdd, nil)
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

func CreateProbeMul(c *gin.Context) {

	pi := []define.ProbeInfoAdd{}

	err := c.ShouldBindJSON(&pi)
	if err != nil {
		log.Error("ShouldBindJSON failed", zap.Error(err))
		resp.JSON(c, resp.ProbeInfoAdd, nil)
		return
	}

	utils.WriteJsonLog(pi)

	for _, v := range pi {
		res := models.GetProbeInfoByName(v.Name)

		if res.Name != "" {
			continue
		}

		err = models.AddProbeInfo(v)
	}

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

func ExportProbeCsv(c *gin.Context) {
	query := ginhelp.GetQueryParams(c)

	data, count := models.GetProbeRes(0, define.ExportLimit, query, "")

	if count == 0 {
		resp.JSON(c, resp.Nodata, nil)
	}

	filename, err := toProbeCsv(data, "ExportProbeCsv")

	if err != nil {
		slog.Println(slog.DEBUG, "t.toCsv() failed == ", err)
	}
	if filename == "" {
		slog.Println(slog.DEBUG, "export excel file failed == ", filename)
	}
	file, err := os.Open("./" + filename)
	if err != nil {
		resp.JSONNew(c, resp.ErrBadRequest, "文件不存在")
		return
	}
	defer file.Close()

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Header("Content-Type", "text/csv;") // Set Content-Type to audio/mpeg
	io.Copy(c.Writer, file)               //直接返回文件

}

func toProbeCsv(data []define.ProbeRes, name string) (string, error) {
	//获取数据

	strTime := time.Now().Format("20060102150405")
	//创建csv文件
	filename := fmt.Sprintf("%s-%s.csv", name, strTime)
	xlsFile, fErr := os.OpenFile("./"+filename, os.O_RDWR|os.O_CREATE, 0766)
	if fErr != nil {
		slog.Println(slog.DEBUG, "Export:created excel file failed ==", fErr)
		return "", fErr
	}
	defer xlsFile.Close()
	//开始写入内容
	//写入UTF-8 BOM,此处如果不写入就会导致写入的汉字乱码
	xlsFile.WriteString("\xEF\xBB\xBF")
	wStr := csv.NewWriter(xlsFile)
	wStr.Write([]string{"ip", "port", "规则名称", "规则分组", "所属国家", "标签", "匹配结果", "是否处理", "备注", "创建时间"})

	for _, s := range data {
		wStr.Write([]string{s.IP, s.Port, s.Pname, s.Pg, s.Region, s.Tags, s.Matched.String(), s.Dealed.String(), s.Remark, s.Ctime.String()})
	}
	wStr.Flush() //写入文件
	return filename, nil
}

func ExportTem(c *gin.Context) {

	filename, err := toCsv("ExportTem")

	if err != nil {
		slog.Println(slog.DEBUG, "t.toCsv() failed == ", err)
	}
	if filename == "" {
		slog.Println(slog.DEBUG, "export excel file failed == ", filename)
	}
	defer func() {
		err := os.Remove("./" + filename) //下载后，删除文件
		if err != nil {
			slog.Println(slog.DEBUG, "remove  excel file failed", err)
		}
	}()
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Writer.Header().Add("Content-Type", "text/csv;charset=gb2312") //设置下载文件格式，流式下载
	c.File("./" + filename)                                          //直接返回文件

}

func toCsv(name string) (string, error) {
	//获取数据

	strTime := time.Now().Format("20060102150405")
	//创建csv文件
	filename := fmt.Sprintf("%s-%s.csv", name, strTime)
	xlsFile, fErr := os.OpenFile("./"+filename, os.O_RDWR|os.O_CREATE, 0766)
	if fErr != nil {
		slog.Println(slog.DEBUG, "Export:created excel file failed ==", fErr)
		return "", fErr
	}
	defer xlsFile.Close()
	//开始写入内容
	//写入UTF-8 BOM,此处如果不写入就会导致写入的汉字乱码
	xlsFile.WriteString("\xEF\xBB\xBF")
	wStr := csv.NewWriter(xlsFile)
	wStr.Write([]string{"probe_desc", "probe_name", "probe_protocol", "probe_send", "probe_recv", "probe_group", "probe_tags", "probe_match_type"})

	wStr.Flush() //写入文件
	return filename, nil
}
