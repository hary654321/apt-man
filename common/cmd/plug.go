package cmd

import (
	"os/exec"
	"strings"
	"zrDispatch/common/utils"
	"zrDispatch/core/slog"
)

func Plug(taskId, ip, port, cmd string) (string, error) {

	ipstr := strings.Join(utils.GetIpArr(ip), " ")

	portstr := strings.Join(utils.GetPortArr(port), ",")

	res := "/tmp/" + taskId + "plugres"

	if !strings.Contains(cmd, "nmap") {
		cmd = "/zrtx/apt/bin"
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
