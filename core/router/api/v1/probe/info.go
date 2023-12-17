package probe

import (
	"bytes"
	"encoding/csv"
	"errors"
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

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"

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

	port := utils.GetPortArr(pi.Port)
	if len(port) < 1 {
		resp.JSONNew(c, resp.ProbeInfoAdd, "请按照规定格式填写port")
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

var mapa = map[string]string{
	"规则名称": "probe_name",
	"规则分组": "probe_group",
	"规则标签": "probe_tags",
	"规则协议": "probe_protocol",
	"匹配类型": "probe_match_type",
	"规则荷载": "probe_send",
	"结果匹配": "probe_recv",
	"描述":   "probe_desc",
	"目标端口": "probe_port",
}

func Import(c *gin.Context) {
	//FormFile返回所提供的表单键的第一个文件
	f, _ := c.FormFile("file")
	//SaveUploadedFile上传表单文件到指定的路径
	c.SaveUploadedFile(f, "./"+f.Filename)

	data, err := GetcsvDataPro(f.Filename, mapa)

	if err != nil {
		resp.JSONNew(c, resp.ErrBadRequest, err.Error())
		return
	}

	res := models.BatchAddProbeInfo(data)

	if res == 0 {
		resp.JSON(c, resp.AddFail, nil)
		return
	}

	resp.JSON(c, resp.Success, map[string]string{"msg": "导入成功"})
}

func GetcsvDataPro(filename string, mapa map[string]string) (ResData []map[string]any, err error) {

	file, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("文件打开失败: ", err)
		return
	}

	reader := csv.NewReader(transform.NewReader(bytes.NewReader(file), simplifiedchinese.GBK.NewDecoder()))
	rowNum := 1
	var headarr []string
	for {
		line, err := reader.Read()
		if err == io.EOF {
			fmt.Println("文件读取完毕")
			break
		}

		if err != nil {
			fmt.Println("读取文件时发生错误: ", err)
			break
		}
		if rowNum == 1 {
			headarr = line
		} else {
			var rowData = make(map[string]any)
			for k, v := range headarr {

				if v == "" {
					continue
				}

				if mapa[v] == "probe_name" {

					res := models.GetProbeInfoByName(line[k])

					if res.Name != "" {
						slog.Println(slog.DEBUG, "名称重复", line[k])
						return ResData, errors.New("有重复名称:" + line[k])
					}

				}

				if mapa[v] == "probe_protocol" && !utils.In_array(line[k], []string{"HTTP", "TCP"}) {

					return ResData, errors.New("协议必须为HTTP,TCP")
				}
				if mapa[v] == "probe_match_type" && !utils.In_array(line[k], []string{"keyword", "==", "re", "cert"}) {

					return ResData, errors.New("匹配类型必须为keyword,re,==,cert")
				}
				rowData[mapa[v]] = line[k]
				// slog.Println(slog.DEBUG, v, "========", line[k])
			}

			ResData = append(ResData, rowData)
		}
		rowNum++

	}
	return
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

	data, count := models.GetProbeInfo(0, define.ExportLimit, query)

	if count == 0 {
		resp.JSON(c, resp.Nodata, nil)
	}

	filename, err := toProbeCsv(data, "规则详情")

	if err != nil {
		slog.Println(slog.DEBUG, "t.toCsv() failed == ", err)
	}
	if filename == "" {
		slog.Println(slog.DEBUG, "export excel file failed == ", filename)
	}

	resp.JSON(c, resp.Success, filename)

}

func toProbeCsv(data []define.ProbeInfoRes, name string) (string, error) {
	//获取数据

	strTime := time.Now().Format("20060102150405")
	//创建csv文件
	filename := fmt.Sprintf("%s-%s.csv", name, strTime)
	xlsFile, fErr := os.OpenFile("tem/"+filename, os.O_RDWR|os.O_CREATE, 0766)
	if fErr != nil {
		slog.Println(slog.DEBUG, "Export:created excel file failed ==", fErr)
		return "", fErr
	}
	defer xlsFile.Close()
	//开始写入内容
	//写入UTF-8 BOM,此处如果不写入就会导致写入的汉字乱码
	xlsFile.WriteString("\xEF\xBB\xBF")
	wStr := csv.NewWriter(xlsFile)
	wStr.Write([]string{"规则名称", "规则分组", "规则标签", "规则协议", "匹配类型", "规则荷载", "结果匹配", "目标端口", "描述"})

	for _, s := range data {
		wStr.Write([]string{s.Name, s.Group, s.Tags, s.Pro, s.MT, s.Send, s.Recv, s.Port, s.Desc})
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
	// wStr.Write([]string{"probe_desc", "probe_name", "probe_protocol", "probe_send", "probe_recv", "probe_group", "probe_tags", "probe_match_type"})
	wStr.Write([]string{"规则名称", "规则分组", "规则标签", "规则协议", "匹配类型", "规则荷载", "结果匹配", "描述"})

	wStr.Flush() //写入文件
	return filename, nil
}
