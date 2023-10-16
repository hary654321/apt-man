package cmd

import (
	"os/exec"
	"strings"
	"zrDispatch/common/utils"
	"zrDispatch/core/slog"
)

func Plug(taskId, ip, port, cmd string) (string, error) {

	if cmd == "" {
		return "", nil
	}

	ipstr := strings.Join(utils.GetIpArr(ip), " ")

	portstr := strings.Join(utils.GetPortArr(port), ",")

	res := "/tmp/" + taskId + "plugres"

	if !strings.Contains(cmd, "nmap") {

		exe := strings.Split(cmd, " ")[0]
		exe = "/app/" + exe
		if !utils.PathExist(exe) {
			slog.Println(slog.DEBUG, "文件不存在", exe)
			return "", nil
		}
		cmd = "/app/" + cmd
		ipstr = strings.Join(utils.GetIpArr(ip), ",")
	} else {
		if !utils.PathExist("/usr/bin/nmap") {
			slog.Println(slog.DEBUG, "nmap不存在")
			return "", nil
		}
	}

	cmd = strings.Replace(cmd, "{ip}", ipstr, -1)

	cmd = strings.Replace(cmd, "{port}", portstr, -1)

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
