package cmd

import (
	"context"
	"os/exec"
	"strings"
	"time"
	"zrDispatch/common/utils"
	"zrDispatch/core/slog"
)

// nmap -T4 -A -v   172.16.130.100-172.16.130.138
func Scan(taskId, ip, port string) (string, error) {

	ipstr := strings.Join(utils.GetIpArr(ip), " ")

	portstr := strings.Join(utils.GetPortArr(port), ",")

	cmdStr := "nmap -T5   -p " + portstr + " -oX /tmp/" + taskId + ".xml " + ipstr

	slog.Println(slog.DEBUG, cmdStr)

	cmd1 := exec.Command("bash", "-c", cmdStr)

	_, err := cmd1.CombinedOutput()
	if err != nil {
		slog.Println(slog.ERROR, err)

		return "", err

	}
	return "/tmp/" + taskId + ".xml", nil
}

func CatRead(file string) string {
	cmd := exec.Command("cat", file)

	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Println(slog.DEBUG, string(out))
		slog.Println(slog.DEBUG, err)
	}

	return string(out)
}

// 获取操作系统信息
func GetOpInfo(ip string) string {
	slog.Println(slog.DEBUG, "开始调用nmap 操作系统：", ip)
	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	cmd := exec.CommandContext(ctx, "nmap", "-O", ip)

	out, err := cmd.CombinedOutput()
	if err != nil {
		//slog.Println(slog.DEBUG,  string(out))
		slog.Println(slog.DEBUG, err)
	}
	//slog.Println(slog.INFO,  string(out))

	res := string(out)
	strArr := strings.Split(res, "\n")

	for _, line := range strArr {
		//slog.Println(slog.DEBUG,  line)
		if utils.GetStrACount("OS details", line) > 0 {
			return SubStrAfter(line, ":")
		}
	}

	elapsed := time.Since(start)
	slog.Println(slog.DEBUG, "该函数执行完成耗时：", elapsed)

	return ""
}

func SubStrAfter(a, needle string) string {
	comma := strings.Index(a, needle)

	return a[comma+1:]

}
