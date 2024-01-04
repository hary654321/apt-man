package main

import (
	"os"
	"zrDispatch/common/log"
	"zrDispatch/core/alarm"
	"zrDispatch/core/client"
	"zrDispatch/core/config"
	"zrDispatch/core/match"
	"zrDispatch/core/model"
	"zrDispatch/core/router"
	"zrDispatch/core/schedule"
	"zrDispatch/core/slog"
	mylog "zrDispatch/core/utils/log"
	"zrDispatch/models"

	"go.uber.org/zap"
)

func main() {

	config.Init(os.Args[1])
	mylog.Init()
	alarm.InitAlarm()
	err := model.InitDb()
	if err != nil {
		log.Fatal("InitDb failed", zap.Error(err))
	}
	model.InitRabc()
	models.Setup()

	go model.Update()
	// model.Pi()

	go client.HeartBeat() // 心跳检测

	go match.Match() // 匹配

	//初始化定时任务
	err = schedule.Init2()
	if err != nil {
		slog.Printf(slog.DEBUG, "init schedule failed", zap.Error(err))
	}

	go client.GetUpdateUserList()

	router.Run()

}
