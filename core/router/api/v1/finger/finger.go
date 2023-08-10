package finger

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
	"zrDispatch/common/utils"
	"zrDispatch/core/alyaze"
	"zrDispatch/core/slog"
	"zrDispatch/e"
	"zrDispatch/models"
)

func GetFinger(c *gin.Context) {
	name := c.Query("name")
	finger := c.Query("finger")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}
	if finger != "" {
		maps["finger"] = finger
	}

	code := e.SUCCESS

	slog.Println(slog.DEBUG, maps)
	data["lists"] = models.GetFinger(utils.GetPage(c), 10, maps)
	data["total"] = models.GetFingerTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddFinger(c *gin.Context) {

	name := c.PostForm("name")
	description := c.PostForm("description")
	finger := c.PostForm("finger")

	code := e.INVALID_PARAMS

	if strings.HasSuffix(finger, ",") {
		code = e.INVALID_FINGER
	} else {

		valid := validation.Validation{}

		// 输入长度限制
		valid.Required(name, "name").Message("名称不能为空")

		if !valid.HasErrors() {

			go func() {
				data := make(map[string]interface{})
				data["name"] = name
				data["description"] = description
				data["finger"] = finger
				models.AddFinger(data)
			}()
			code = e.SUCCESS

		}

	}

	c.JSON(http.StatusOK, gin.H{
		"code": 400,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

// 新增指纹时测试指纹是否正确
func TestFinger(c *gin.Context) {

	testurl := c.PostForm("testurl")
	finger := c.PostForm("finger")

	valid := validation.Validation{}

	// 输入长度限制
	valid.Required(finger, "finger").Message("finger不能为空")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if CheckFingerSingle(testurl, finger) {
			code = e.SUCCESS
		} else {
			code = 402
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

// 扫描当前指纹在数据库中的数量
func ScanFinger(c *gin.Context) {

	id := com.StrTo(c.Param("id")).MustInt()
	code := e.SUCCESS
	CheckFingerOne(id)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

func DeleteFinger(c *gin.Context) {

	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Required(id, "id").Message("id不能为空")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {

		//删除任务结果记录
		models.DeleteFinger(id)
		code = e.SUCCESS

	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

var (
	wa     *alyaze.WebAnalyzer
	thread = 5
	err    error

	crawlCount      = 0
	searchSubdomain = false
	redirect        = false
)

// 测试单个指纹是否编写正常
func CheckFingerSingle(domain, finger string) bool {

	var allFingers bytes.Buffer
	allFingers.WriteString(`{"technologies": {`)

	allFingers.WriteString(finger)
	Fingers := strings.TrimSuffix(allFingers.String(), ",")
	Fingers = Fingers + " }}"
	//fmt.Println("Fingers:",Fingers)

	if wa, err = alyaze.NewWebAnalyzer(Fingers, nil); err != nil {
		log.Printf("initialization failed: %v", err)
		return false
	}

	job := alyaze.NewOnlineJob(domain, "", nil, crawlCount, searchSubdomain, redirect)
	result, _ := wa.Process(job)
	if len(result.Matches) > 0 {
		return true
	}
	return false

}

// 扫描数据库资产在指纹的结果
func CheckFingerOne(id int) {
	var success int

	start := time.Now()
	allFinger := models.GetAllFingerId(id)
	HttpRes := models.GetIplistHttp()

	costTime := time.Since(start)

	data := make(map[string]interface{})
	data["taskid"] = 0
	data["task_name"] = "checkFinger"
	data["task_type"] = "checkFinger"
	data["all_num"] = 0
	data["succes_num"] = 0
	data["run_time"] = ""
	data["error"] = "nil"
	data["status"] = 0
	//taskId := models.AddLog(data)

	domains := make(chan string)

	var allFingers bytes.Buffer
	allFingers.WriteString(`{"technologies": {`)

	for _, finger := range allFinger {
		allFingers.WriteString(finger.Finger)
		allFingers.WriteString(",")
	}
	Fingers := strings.TrimSuffix(allFingers.String(), ",")
	Fingers = Fingers + " }}"
	//fmt.Println("Fingers:",Fingers)

	if wa, err = alyaze.NewWebAnalyzer(Fingers, nil); err != nil {
		log.Printf("initialization failed: %v", err)
	}

	var wg sync.WaitGroup
	for i := 0; i < thread; i++ {
		wg.Add(1)
		go func() {

			for host := range domains {
				fmt.Println("now is :", host)
				job := alyaze.NewOnlineJob(host, "", nil, crawlCount, searchSubdomain, redirect)
				result, links := wa.Process(job)

				if searchSubdomain {
					for _, v := range links {
						crawlJob := alyaze.NewOnlineJob(v, "", nil, 0, false, redirect)
						wa.Process(crawlJob)
					}
				}

				for _, a := range result.Matches {
					success++
					dataUpdate := make(map[string]interface{})
					dataUpdate["cms"] = a.AppName
					models.EditIplistByUrl(result.Host, dataUpdate)
				}

			}

			wg.Done()
		}()
	}
	for _, k := range HttpRes {
		if k.Loginurl != "" {
			domains <- k.Loginurl
		}
	}

	close(domains)
	wg.Wait()

	data = make(map[string]interface{})
	data["all_num"] = success
	data["status"] = 1
	data["run_time"] = fmt.Sprintf("%s", costTime)
	//models.EditLog(taskId, data)

}
