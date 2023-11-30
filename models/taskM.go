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

	slog.Println(slog.DEBUG, pge.Run)

	res := db.Table("task").Select("TaskType", "ip", "port", "cronExpr", "timeout", "routePolicy", "remark", "probeId", "plug", "hostGroupID", "run", "threads", "name", "group", "priority").Model(&define.ChangeTask{}).Where("id = ?", pge.ID).Updates(pge)

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

func ChangeTaskRun(id string, run int) error {

	res := db.Table("task").Where("id", id).Update("run", run)

	return res.Error
}

func ChangeTaskStatus(id string, run define.TaskOneStatus) error {

	res := db.Table("task").Where("id", id).Update("status", run)

	return res.Error
}

func UpdateTaskCreate(id, uid string) {

	db.Table("task").Where("id = ?", id).Update("createByID", uid)
}

func GetTaskInfoByName(name string) (ProbeMatch define.CreateTask) {
	dbTmp := db.Table("task")

	dbTmp.Where("name = ?", name).Where("isDeleted", 0).Take(&ProbeMatch)

	return
}

func GetTaskMap(ids []string) map[string]define.GetIdNameGroup {
	dbTmp := db.Table("task")

	var ProbeInfo []define.GetIdNameGroup

	dbTmp.Where("id in ?", ids).Find(&ProbeInfo)

	pgmap := make(map[string]define.GetIdNameGroup)
	for _, v := range ProbeInfo {
		pgmap[v.ID] = v
	}

	return pgmap
}

func GetTaskIds(name string) (res []string) {
	dbTmp := db.Table("task")

	var ProbeInfo []define.GetIdNameGroup
	dbTmp.Where("name like ?", name+"%").Find(&ProbeInfo)

	for _, v := range ProbeInfo {
		res = append(res, v.ID)
	}
	return
}

func GetTaskIdsBygrop(name string) (res []string) {
	dbTmp := db.Table("task")

	var ProbeInfo []define.GetIdNameGroup
	dbTmp.Where("`group` like ?", name+"%").Find(&ProbeInfo)

	for _, v := range ProbeInfo {
		res = append(res, v.ID)
	}
	return
}
