package res

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
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
	data, count := models.GetProbeRes(q.Offset, q.Limit, query)

	resp.JSON(c, resp.Success, data, int(count))
}

func ExportProbeCsv(c *gin.Context) {
	query := ginhelp.GetQueryParams(c)

	data, count := models.GetProbeRes(0, define.ExportLimit, query)

	if count == 0 {
		resp.JSON(c, resp.Nodata, nil)
	}

	filename, err := toProbeCsv(data, "匹配结果")

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
	c.Writer.Header().Add("Content-Type", "application/octet-stream") //设置下载文件格式，流式下载
	c.File("./" + filename)                                           //直接返回文件

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
	wStr.Write([]string{"ip", "port", "探针名称", "探针分组", "所属国家", "标签", "匹配结果", "创建时间"})

	for _, s := range data {
		wStr.Write([]string{s.IP, s.Port, s.Pname, s.Pg, s.Region, s.Tags, s.Matched.String(), s.Ctime.String()})
	}
	wStr.Flush() //写入文件
	return filename, nil
}
