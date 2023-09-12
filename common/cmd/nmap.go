package cmd

import (
	"os/exec"
	"strings"
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

func Plug(taskId, ip, port, cmd string) (string, error) {

	ipstr := strings.Join(utils.GetIpArr(ip), " ")

	portstr := strings.Join(utils.GetPortArr(port), ",")

	res := "/tmp/" + taskId + ".xml"

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
