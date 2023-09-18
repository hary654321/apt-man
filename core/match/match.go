package match

import (
	"regexp"
	"strings"
	"time"
	"zrDispatch/core/slog"
	"zrDispatch/core/utils/define"
	"zrDispatch/models"
)

// 匹配
func Match() {
	sleepTime := time.Duration(1 * time.Minute)
	for {
		res := models.GetNotMacthedList()

		if len(res) > 0 {
			sleepTime = time.Duration(1 * time.Second)
		} else {
			sleepTime = time.Duration(1 * time.Minute)
		}

		for _, pres := range res {
			Macth(pres)
		}
		time.Sleep(sleepTime)
		slog.Println(slog.DEBUG, "匹配工作进行中...")
	}
}

func Macth(ps define.ProbeRes) {
	probeInfo := models.GetProbeInfoByName(ps.Pname)

	if probeInfo.MT == "keyword" {
		if strings.Contains(ps.Res, probeInfo.Recv) {
			models.UpdateProbeMatch(ps.Id, define.Matched)
			return
		}
	}

	if probeInfo.MT == "re" {

		match, err := regexp.MatchString(probeInfo.Recv, ps.Res)

		slog.Println(slog.DEBUG, "正则匹配结果：", match, err)

		if match {
			models.UpdateProbeMatch(ps.Id, define.Matched)
			return
		}
	}

	if probeInfo.MT == "==" {
		if probeInfo.Recv == ps.Res {
			models.UpdateProbeMatch(ps.Id, define.Matched)
			return
		}
	}

	if probeInfo.MT == "cert" {
		if strings.Contains(ps.Cert, probeInfo.Recv) {
			models.UpdateProbeMatch(ps.Id, define.Matched)
			return
		}
	}

	models.UpdateProbeMatch(ps.Id, define.NotMatched)

}
