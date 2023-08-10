package models

import "zrDispatch/core/utils/define"

// 保存结果
func AddPortRes(c define.PortScan) int {

	res := db.Table("port_result").Create(&c)
	return int(res.RowsAffected)
}

func GetPortRes(pageNum int, pageSize int, maps map[string]interface{}) (PortRes []define.PortRes, total int64) {
	dbTmp := db.Table("port_result")

	dbTmp.Where(maps).Count(&total)

	dbTmp.Where(maps).Offset(pageNum).Limit(pageSize).Order("id  desc").Find(&PortRes)

	return
}

func DeletePortRes(ids []int) int64 {

	res := db.Table("port_result").Where("probe_id in (?) ", ids).Delete(&define.PortRes{})

	return res.RowsAffected
}

func GetTaskPortRes(taskId string) (PortRes []define.PortRes) {
	dbTmp := db.Table("port_result")

	dbTmp.Where("run_task_id like ? ", taskId+"%").Order("id  desc").Find(&PortRes)

	return
}

func GetTaskLiveIpCount(taskId string) (ipcount int64) {
	dbTmp := db.Table("port_result")

	dbTmp.Where("run_task_id like ? ", taskId+"%").Select("distinct ip").Distinct("ip").Count(&ipcount)

	return
}

func GetOsSelect() (os []define.Os) {
	dbTmp := db.Table("port_result")

	dbTmp.Where("os!=?", "").Select("os").Distinct("os").Find(&os)

	return
}
