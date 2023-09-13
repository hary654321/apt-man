package cmd

import (
	"os/exec"
	"zrDispatch/common/utils"
	"zrDispatch/core/slog"
)

func RunLog() string {

	path := utils.GetCurrentDirectory()

	cmd := exec.Command("tail", "-100", path+"/m.log")

	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}

	return string(out)
}

func Exec(cmdStr string) (string, error) {
	cmd := exec.Command("bash", "-c", cmdStr)

	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Println(slog.DEBUG, cmdStr, "=======", string(out), "====", err)
		return string(out), err
	}

	return string(out), err
}
