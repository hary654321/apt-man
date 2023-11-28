package models

import (
	"zrDispatch/common/utils"
	"zrDispatch/core/slog"
	"zrDispatch/core/utils/define"
)

// 创建任务
func AddProbeRes(c define.ProbeResCreate) int {

	res := db.Table("probe_result").Create(&c)
	return int(res.RowsAffected)
}

func GetProbeRes(pageNum int, pageSize int, maps map[string]interface{}, order string) (ProbeRes []define.ProbeRes, total int64) {

	if order != "" {
		order = "probe_result." + order
	} else {
		order = "probe_result.id  desc"
	}

	slog.Println(slog.DEBUG, maps)

	dbTmp := db.Table("probe_result")

	dbTmp = dbTmp.Select("task.name as task_name,task.group as task_group,os.os,probe_result.create_time,probe_result.id,probe_result.ip,probe_result.run_task_id,probe_result.port,probe_result.probe_name,probe_result.cert,probe_result.matched,probe_result.response,probe_result.dealed,probe_result.remark, probe_info.probe_send,probe_info.probe_recv,probe_info.probe_group,probe_info.probe_tags").
		Joins("left join probe_info on probe_info.probe_name = probe_result.probe_name").
		Joins("left join os on probe_result.ip = os.ip").
		Joins("left join task on task.id = probe_result.task_id")

	if utils.GetInterfaceToString(maps["probe_group"]) != "" {
		dbTmp = dbTmp.Where("probe_info.probe_group = ?", utils.GetInterfaceToString(maps["probe_group"]))
		delete(maps, "probe_group")
	}

	if maps["probe_name"] != "" {
		dbTmp = dbTmp.Where("probe_result.probe_name LIKE ?", "%"+utils.GetInterfaceToString(maps["probe_name"])+"%")
		delete(maps, "probe_name")
	}

	if maps["task_name"] != "" {
		dbTmp = dbTmp.Where("task.name LIKE ?", "%"+utils.GetInterfaceToString(maps["task_name"])+"%")
		delete(maps, "task_name")
	}

	if maps["task_group"] != "" {
		dbTmp = dbTmp.Where("task.group LIKE ?", "%"+utils.GetInterfaceToString(maps["task_group"])+"%")
		delete(maps, "task_group")
	}

	dbTmp.Where(maps).Count(&total)

	dbTmp.Where(maps).Offset(pageNum).Limit(pageSize).Order(order).Find(&ProbeRes)

	var ProbeResNew []define.ProbeRes

	pgMap := GetPgMap()
	for _, v := range ProbeRes {
		v.Region = pgMap[v.Pg]
		ProbeResNew = append(ProbeResNew, v)
	}

	return ProbeResNew, total
}

func GetNotMacthedList() (ProbeRes []define.ProbeRes) {
	dbTmp := db.Table("probe_result")

	dbTmp.Where("matched", define.MatchInit).Limit(1000).Order("id  asc").Find(&ProbeRes)

	return
}

func DeleteProbeRes(ids []int) int64 {

	res := db.Table("probe_result").Where("probe_id in (?) ", ids).Delete(&define.ProbeRes{})

	return res.RowsAffected
}

func UpdateProbeMatch(id int, matched define.MatchStatus) error {

	res := db.Table("probe_result").Where("id = ?", id).Update("matched", matched)

	return res.Error
}

func GetTaskProbe(taskId string) []define.ProbeResJJ {
	dbTmp := db.Table("probe_result")

	var ProbeRes []define.ProbeResJJ
	dbTmp.Select("ip,port,probe_name").Where("task_id = ? ", taskId).Where("matched", define.Matched).Order("id  desc").Find(&ProbeRes)

	havemap := make(map[string]int)

	var ProbeResUnique []define.ProbeResJJ
	for _, v := range ProbeRes {

		if havemap[v.IP+v.Port] == 1 {
			continue
		}

		ProbeResUnique = append(ProbeResUnique, v)

		havemap[v.IP+v.Port] = 1
	}

	return ProbeResUnique
}

func GetTaskMatchIpCount(taskId string) (ipcount int64) {
	dbTmp := db.Table("probe_result")

	dbTmp.Where("run_task_id like ? ", taskId+"%").Where("matched", define.Matched).Select("distinct ip").Distinct("ip").Count(&ipcount)

	return
}

// 通过id，更新
func EditProbeRes(pge define.ProbeResEdit) int64 {

	res := db.Table("probe_result").Where("id = ?", pge.Id).Updates(pge)
	slog.Println(slog.DEBUG, res.Error)
	return res.RowsAffected
}
