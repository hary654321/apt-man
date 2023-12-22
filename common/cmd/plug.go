package cmd

import (
	"os"
	"os/exec"
	"strings"
	"zrDispatch/common/utils"
	"zrDispatch/core/slog"
)

func Plug(taskId, ip, plugname, cmd string) (string, error) {

	if cmd == "" {
		return "", nil
	}

	ipstr := strings.Join(utils.GetIpArr(ip), " ")

	dir, _ := os.Getwd()

	path := dir + "/plugres/" + plugname

	if !utils.PathExist(path) {
		err := os.MkdirAll(path, 0777)
		if err != nil {
			slog.Println(slog.WARN, "新建日志目录失败")
		}
	}

	res := path + "/" + taskId + ".res"

	if !strings.Contains(cmd, "nmap") {

		exe := strings.Split(cmd, " ")[0]
		exe = dir + "/" + exe
		if !utils.PathExist(exe) {
			slog.Println(slog.DEBUG, "文件不存在", exe)
			return "", nil
		}
		cmd = dir + "/" + cmd
		ipstr = strings.Join(utils.GetIpArr(ip), ",")
	} else {
		if !utils.PathExist("/usr/bin/nmap") {
			slog.Println(slog.DEBUG, "nmap不存在")
			return "", nil
		}
	}

	cmd = strings.Replace(cmd, "{ip}", ipstr, -1)

	cmd = strings.Replace(cmd, "{res}", res, -1)

	slog.Println(slog.DEBUG, cmd)

	cmd1 := exec.Command("bash", "-c", cmd)

	_, err := cmd1.CombinedOutput()
	if err != nil {
		slog.Println(slog.ERROR, err)

		return "", err

	}
	return res, nil
}
