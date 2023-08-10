package cmd

import (
	"os/exec"
	"strconv"
	"strings"
	"zrDispatch/common/utils"
	"zrDispatch/core/config"
	"zrDispatch/core/slog"
)

func CmdIpCount(task_id, date string) int {

	ipLastPath := task_id + ".txt"
	if date != "" {
		ipLastPath = date + "|" + ipLastPath
	}

	slog.Println(slog.DEBUG, ipLastPath)

	path := utils.GetCurrentDirectory()

	cmd := exec.Command("wc", "-l", path+"/"+ipLastPath)

	slog.Println(slog.DEBUG, path+"/"+ipLastPath)

	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}

	resArr := strings.Split(string(out), " ")

	num, _ := strconv.Atoi(resArr[0])

	return num
}

func ResCount(hostname, date, file string) int {

	path := config.CoreConf.Server.LogPath + date + "/" + hostname + "/cyberspace/" + file + date + ".json"

	slog.Println(slog.DEBUG, path)

	cmd := exec.Command("wc", "-l", path)

	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}

	resArr := strings.Split(string(out), " ")

	num, _ := strconv.Atoi(resArr[0])

	return num

}

func RestartTime(hostname, date string) int {

	path := config.CoreConf.Server.LogPath + date + "/" + hostname + "/cyberspace/" + "restart" + date + ".json"

	slog.Println(slog.DEBUG, path)

	cmd := exec.Command("wc", "-l", path)

	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}

	resArr := strings.Split(string(out), " ")

	num, _ := strconv.Atoi(resArr[0])

	return num

}

func ResSize(hostname, date, file string) string {

	path := config.CoreConf.Server.LogPath + date + "/" + hostname + "/cyberspace/" + file + date + ".json"

	cmd := exec.Command("du", "-sh", path)

	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}
	slog.Println(slog.DEBUG, path, "-----------", string(out))
	resArr := strings.Split(string(out), "/u2")

	return resArr[0]

}

func RunLog() string {

	path := utils.GetCurrentDirectory()

	cmd := exec.Command("tail", "-100", path+"/m.log")

	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}

	return string(out)
}

func Restart() string {

	cmd := exec.Command("bash", "-c", "ps -ef | grep ./manage | grep -v grep | awk '{print $2}' | xargs kill -9")

	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}

	cmd1 := exec.Command("bash", "-c", "cd /u2/cy/www/strategy-manage && ./manage 138.toml>> m.log &")

	cmd1.CombinedOutput()
	return string(out)
}

func ClearLog(line string) {
	cmd := exec.Command("bash", "-c", "sed -i '1,"+line+"d' /var/www/cyberspacemapping/www/strategy-manage/m.log")

	cmd.CombinedOutput()
}

func IpInfoFile() string {
	cmd := exec.Command("bash", "-c", "tree /u2/zrtx/log/cyberspace |grep res")

	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}

	cmd1 := exec.Command("bash", "-c", "cd /var/www/cyberspacemapping/www/strategy-manage && ./manage>> m.log &")

	cmd1.CombinedOutput()
	return string(out)
}
