package models

import "zrDispatch/core/utils/define"

func GetOneLog(runTaskId string) (log define.Log) {
	db.Table("log").Where("runTaskId = ? ", runTaskId).Take(&log)
	return
}

func UpdatePress(runTaskId string, p int) error {
	res := db.Table("log").Where("runTaskId = ? ", runTaskId).Where("progress <= ?", p).Update("progress", p)
	return res.Error
}

func UpdateResReason(runTaskId string, status int, err string, endtime string) error {
	res := db.Table("log").Where("runTaskId = ? ", runTaskId).Update("taskresps", err).Update("status", status).Update("endtime", endtime)
	return res.Error
}
