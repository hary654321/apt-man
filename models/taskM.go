package models

import (
	"zrDispatch/core/slog"
	"zrDispatch/core/utils/define"
)

func AddTask(data define.CreateTask) error {

	if data.Port == "" {
		data.Port = "22,80,443"
	}

	res := db.Table("task").Create(data)
	return res.Error

}

func EditTask(pge define.ChangeTask) error {

	res := db.Table("task").Select("TaskType", "ip", "port", "cronExpr", "timeout", "routePolicy", "remark", "probeId", "hostGroupID", "run", "threads", "name").Model(&define.ChangeTask{}).Where("id = ?", pge.ID).Updates(pge)

	return res.Error
}

func GetTaskByID(id string) (task *define.DetailTask, err error) {
	res := db.Table("task").Where("id = ? ", id).Take(&task)
	err = res.Error
	slog.Println(slog.DEBUG, task)
	return
}

func GetTaskByIDClone(id string) (task *define.CreateTask, err error) {
	res := db.Table("task").Where("id = ? ", id).Take(&task)
	err = res.Error
	slog.Println(slog.DEBUG, task)
	return
}

func DeleteTask(id string) error {

	res := db.Table("task").Where("id", id).Update("isDeleted", 1)

	return res.Error
}
