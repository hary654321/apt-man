package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"zrDispatch/core/config"
	"zrDispatch/core/slog"

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// 返回一个32位md5加密后的字符串
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 获取系统当前日期
func GetDate() string {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	return time.Now().In(cstZone).Format("2006-01-02")
}

// 获取系统当前日期
func GetTimeStr() string {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	return time.Now().In(cstZone).Format("2006-01-02 15:04:05")
}

// 前一天
func GetLastDate() string {
	now := time.Now().AddDate(0, 0, -1) //获取当前时间对象
	year := strconv.Itoa(now.Year())
	month := now.Format("01")
	day := now.Format("02")
	//rangeDate := []string{"18", "19", "20", "21", "22", "23"}

	return year + "-" + month + "-" + day
}

// 前一小时
func GetLastHour() string {
	var cstZone = time.FixedZone("CST", 7*3600) // 东八
	return time.Now().In(cstZone).Format("2006-01-02-15")
}

func GetHour() string {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	return time.Now().In(cstZone).Format("2006-01-02-15")
}

func GetTimeStrHIS(t int64) string {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八

	return time.Unix(t, 0).In(cstZone).Format("2006-01-02 15:04:05")
}

func RandInt(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return r.Intn(max)
}

func ArrayToString(arr []string) string {
	var result string
	for _, i := range arr { //遍历数组中所有元素追加成string
		result += i + ","
	}
	return result
}

func MergeMap(x, y map[string]string) map[string]string {

	n := make(map[string]string)
	for i, v := range x {
		for j, w := range y {
			if i == j {
				n[i] = w

			} else {
				if _, ok := n[i]; !ok {
					n[i] = v
				}
				if _, ok := n[j]; !ok {
					n[j] = w
				}
			}
		}
	}

	return n
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// BindArgsWithGin 绑定请求参数
func BindArgsWithGin(c *gin.Context, req interface{}) error {
	return c.ShouldBindWith(req, binding.Default(c.Request.Method, c.ContentType()))
}

// MakeMD5 MD5加密
func MakeMD5(data string) string {
	h := md5.New()
	h.Write([]byte(data)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr) // 输出加密结果
}

// Random 生成随机数
func Random(min, max int) int {
	if min == max {
		return max
	}
	max = max + 1
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return min + r.Intn(max-min)
}

// RandomStr 随机字符串
func RandomStr(l int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	seed := "1234567890QWERTYUIOPASDFGHJKLZXCVBNM"
	str := ""
	length := len(seed)
	for i := 0; i < l; i++ {
		point := r.Intn(length)
		str = str + seed[point:point+1]
	}
	return str
}

// BuildPassword 构建用户密码
func BuildPassword(password, salt string) string {
	return MakeMD5(password + salt)
}

// TernaryOperation 三元操作符
func TernaryOperation(exist bool, res, el interface{}) interface{} {
	if exist {
		return res
	}
	return el
}

// GetBeforeDate 获取n天前的时间
func GetDateFromNow(n int) time.Time {
	timer, _ := time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), time.Local)
	if n == 0 {
		return timer
	}
	return timer.AddDate(0, 0, n)
}

// StrArrExist 检测string数组中是否包含某个字符串
func StrArrExist(arr []string, check string) bool {
	for _, v := range arr {
		if v == check {
			return true
		}
	}
	return false
}

// RetryFunc 带重试的func
func RetryFunc(times int, f func() error) error {
	var (
		reTimes int
		err     error
	)
RETRY:
	if err = f(); err != nil {
		if reTimes == times {
			return err
		}
		time.Sleep(time.Duration(1) * time.Second)
		reTimes++
		goto RETRY
	}
	return nil
}

func GetNowTome() string {
	return time.Now().Format("20060102150405")
}

func GetRandomString(l int) string {
	str := "abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func WirteFileAppend(fileName string, newSubDomain []string) {
	fd, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	for _, st := range newSubDomain {
		buf := []byte(st + "\n")
		fd.Write(buf)
	}
	fd.Close()
}

func PathExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func SliceToString(slices []string) (result string) {
	b, err := json.Marshal(slices)
	if err != nil {
		return
	}
	result = string(b)
	return
}

func SliceSToString(slices [][]string) (result string) {
	b, err := json.Marshal(slices)
	if err != nil {
		return
	}
	result = string(b)
	return
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func WriteZip(day, name string, buf []byte) {
	//slog.Println(slog.WARN, path)
	logPath := config.CoreConf.Server.LogPath + day
	_, err := os.Stat(logPath)
	if err != nil {
		os.MkdirAll(logPath, 0777)
	}

	f, err := os.OpenFile(logPath+"/"+name, os.O_CREATE+os.O_RDWR, 0664)
	if err != nil {
		return
	}

	f.Write(buf)

	//slog.Println(slog.DEBUG, "图片写入完成")
}

func RanNum(max int) int {

	//slog.Println(slog.INFO, max)
	rand.Seed(time.Now().UnixNano())

	// 表示生成 [0,50)之间的随机数
	res := rand.Intn(max)

	//slog.Println(slog.INFO, res)
	return res
}

func RanStr(max int) string {

	//slog.Println(slog.INFO, max)
	rand.Seed(time.Now().UnixNano())

	// 表示生成 [0,50)之间的随机数
	res := rand.Intn(max)

	//slog.Println(slog.INFO, res)
	return GetInterfaceToString(res)
}

func GetCurrentDirectory() string {
	//返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		slog.Println(slog.WARN, err)
	}
	//将\替换成/
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}

func execPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	re, err := filepath.Abs(file)
	if err != nil {
		logs.Error("The eacePath Failed: %s\n", err.Error())
	}
	slog.Println(slog.WARN, re)
	return filepath.Abs(file)
}

func String2int64(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		slog.Println(slog.WARN, err)
	}
	return i
}
