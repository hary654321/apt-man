package client

import (
	"context"
	"time"
	"zrDispatch/common/utils"
	"zrDispatch/core/config"
	"zrDispatch/core/model"
	redis2 "zrDispatch/core/redis"
	"zrDispatch/core/slog"
	sshclient "zrDispatch/core/ssh"
	"zrDispatch/core/utils/define"

	"go.uber.org/zap"
)

func HeartBeat() {

	var restart = "restart.log"
	utils.WriteAppend(restart, utils.GetTimeStr())
	for {
		ctx, cancel := context.WithTimeout(context.Background(),
			config.CoreConf.Server.DB.MaxQueryTime.Duration)
		defer cancel()
		hosts, count, err := model.GetHostsWithStatus(ctx, ">=", model.Deployed)

		if err != nil {
			slog.Println(slog.DEBUG, "BindQuery offset failed", zap.Error(err))
		}
		if count < 1 {
			slog.Println(slog.DEBUG, "无存活主机")
		}
		for _, hostInfo := range hosts {
			// slog.Println(slog.DEBUG, "callWorker", hostInfo.Ip)
			go callWorker(hostInfo)
		}
		slog.Println(slog.DEBUG, "心跳检测进行中")
		time.Sleep(15 * time.Second)
	}
}

func callWorker(hostInfo *define.Host) {
	redisClient := redis2.GetClient()
	dieKey := hostInfo.Ip + "dieCount"
	req := RequestHeartBeat(hostInfo.Ip, hostInfo.ServicePort)
	// slog.Println(slog.DEBUG, hostInfo.Ip, "心跳检测", req)
	if req != nil {
		redisClient.Del(dieKey)
		ctx, cancel := context.WithTimeout(context.Background(),
			config.CoreConf.Server.DB.MaxQueryTime.Duration)
		defer cancel()
		// slog.Println(slog.DEBUG, "UpdateHostHearbeat", hostInfo.Ip)
		err := model.UpdateHostHearbeat(ctx, hostInfo.Ip, hostInfo.ServicePort, utils.GetInterfaceToString(req["time"]), utils.GetInterfaceToString(req["version"]), utils.GetInterfaceToString(req["runningTasks"]))
		if err != nil {
			slog.Println(slog.DEBUG, "UpdateHostHearbeat failed", err)
		}
		// slog.Println(slog.DEBUG, hostInfo.Ip+": pong:"+utils.GetInterfaceToString(req["time"]))
	} else {

		redisClient.Incr(dieKey)
		dieCount, _ := redisClient.Get(dieKey).Int()
		slog.Println(slog.DEBUG, hostInfo.Ip+"dieCount", dieCount)
		if dieCount > 50 {
			redisClient.Del(dieKey)
			slog.Println(slog.DEBUG, hostInfo.Ip+"下线了重启")
			sshclient.Restart(hostInfo)
		}
	}

}
