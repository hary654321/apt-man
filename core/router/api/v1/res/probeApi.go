package res

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"time"
	"zrDispatch/common/utils"
	"zrDispatch/core/ginhelp"
	"zrDispatch/core/slog"
	"zrDispatch/core/utils/define"
	"zrDispatch/core/utils/resp"
	"zrDispatch/models"

	"github.com/gin-gonic/gin"
)

func GetProbeRes(c *gin.Context) {

	query := ginhelp.GetQueryParams(c)

	var q define.Query

	err := c.BindQuery(&q)
	if err != nil {
		slog.Printf(slog.DEBUG, "BindQuery offset failed", err)
	}

	if q.Limit == 0 {
		q.Limit = define.DefaultLimit
	}
	// slog.Println(slog.DEBUG, maps)
	data, count := models.GetProbeRes(q.Offset, q.Limit, query, q.Order)

	resp.JSON(c, resp.Success, data, int(count))
}

func UpdateRemark(c *gin.Context) {

	pg := define.ProbeResEdit{}

	err := c.ShouldBindJSON(&pg)
	if err != nil {
		slog.Println(slog.DEBUG, "ShouldBindJSON failed", err)
		resp.JSON(c, resp.ErrBadRequest, nil)
		return
	}

	data := make(map[string]interface{})

	pg.Utime = utils.GetTimeStr()
	data["count"] = models.EditProbeRes(pg)

	code := resp.Success
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
	})
}

func ExportProbeCsv(c *gin.Context) {
	query := ginhelp.GetQueryParams(c)

	data, count := models.GetProbeRes(0, define.ExportLimit, query, "")

	if count == 0 {
		resp.JSON(c, resp.Nodata, nil)
	}

	filename, err := toProbeCsv(data, "规则扫描结果")

	if err != nil {
		slog.Println(slog.DEBUG, "t.toCsv() failed == ", err)
	}
	if filename == "" {
		slog.Println(slog.DEBUG, "export excel file failed == ", filename)
	}
	resp.JSON(c, resp.Success, filename)

}

func toProbeCsv(data []define.ProbeRes, name string) (string, error) {
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
	wStr.Write([]string{"ip", "port", "威胁名称", "威胁类型", "所属国家", "威胁标签", "匹配结果", "是否处理", "备注", "创建时间"})

	for _, s := range data {
		wStr.Write([]string{s.IP, s.Port, s.Pname, s.Pg, s.Region, s.Tags, s.Matched.String(), s.Dealed.String(), s.Remark, s.Ctime.String()})
	}
	wStr.Flush() //写入文件
	return filename, nil
}
