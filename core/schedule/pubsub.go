package schedule

import (
	"context"
	"encoding/json"
	"time"

	"zrDispatch/core/slog"
	"zrDispatch/core/utils/define"
	"zrDispatch/models"

	"zrDispatch/common/log"
	"zrDispatch/core/config"
	"zrDispatch/core/model"

	"go.uber.org/zap"
)

// TaskEvent task event
type TaskEvent uint8

const (
	// AddEvent recv add event
	AddEvent TaskEvent = iota + 1
	// ChangeEvent recv delete task
	ChangeEvent
	// DeleteEvent recv delete task
	DeleteEvent
	// RunEvent run a task
	RunEvent
	// KillEvent recv stop task
	KillEvent
)

const (
	pubsubChannel = "task.event"
)

// EventData sub data from redis
// 应用于调度节点集群，当添加任务、删除修改任务、终止任务时，
// 所有的集群调度节点都会接收到信息，然后进行相应的修改操作
type EventData struct {
	TaskID string    // task id
	TE     TaskEvent // task event: add change delete stop task
}

// RecvEvent recv task event
func RecvEvent() {
	sub := Cron2.redis.Subscribe(pubsubChannel)
	for msg := range sub.Channel() {
		log.Debug("recv event", zap.String("data", msg.Payload))
		go dealEvent([]byte(msg.Payload))
	}
}

func dealEvent(data []byte) {
	var subdata EventData
	err := json.Unmarshal(data, &subdata)
	if err != nil {
		log.Error("json.Unmarshal event data failed", zap.Error(err))
		return
	}
	switch subdata.TE {
	case AddEvent:
		fallthrough
	case ChangeEvent:
		ctx, cancel := context.WithTimeout(context.Background(), config.CoreConf.Server.DB.MaxQueryTime.Duration)
		defer cancel()
		task, err := model.GetTaskByID(ctx, subdata.TaskID)
		if err != nil {
			log.Error("model.GetTaskByID failed", zap.Error(err))
			return
		}
		slog.Println(slog.WARN, "ChangeEvent",Cron2.ts)

		time.Sleep(time.Duration((3100-task.Priority*1000)*2) * time.Millisecond)

		if task.Cronexpr == "" {
			loop := true
			for loop {

				taskInfo, err := models.GetTaskByID(subdata.TaskID)
				if err != err {
					slog.Println(slog.DEBUG, err)
					return
				}
				slog.Println(slog.DEBUG, subdata, "=======", taskInfo)
				if taskInfo.Status == define.TASK_STATUS_STOP {
					slog.Println(slog.DEBUG, taskInfo.Name, "被终止了")
					return
				}

				loop = false
				for _, ot := range Cron2.ts {
					//有优先级高的
					if ot.status == define.TASK_STATUS_RUNING && ot.Priority > task.Priority && ot.cronexpr == "" {
						slog.Println(slog.DEBUG, task.Name, "====等待===", ot.name, "执行中")
						loop = true
						time.Sleep(3 * time.Second)
						break
					}
				}
			}
		}

		Cron2.addtask(task.ID, task.Name, task.Cronexpr, GetRoutePolicy(task.HostGroupID, task.RoutePolicy), task.Run, task.Status, task.Priority)
	case DeleteEvent:
		Cron2.deletetask(subdata.TaskID)
	case RunEvent:
		task, ok := Cron2.GetTask(subdata.TaskID)
		if !ok {
			log.Error("Can not get Task", zap.String("taskid", subdata.TaskID))
			return
		}
		go task.StartRun(define.Manual)
	case KillEvent:
		Cron2.killtask(subdata.TaskID)
	default:
		log.Warn("unsupport task event", zap.Any("data", subdata))
	}
}

