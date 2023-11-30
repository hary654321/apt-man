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

	dbTmp = dbTmp.Select("*")

	if utils.GetInterfaceToString(maps["probe_group"]) != "" {
		pnames := GetPname(utils.GetInterfaceToString(maps["probe_group"]))
		dbTmp = dbTmp.Where("probe_result.probe_name  in ?", pnames)
		delete(maps, "probe_group")
	}

	if maps["os"] != nil {
		slog.Println(slog.DEBUG, "os", maps["os"])
		dbTmp = dbTmp.Where("probe_result.ip in ?", GetOsip(utils.GetInterfaceToString(maps["os"])))
		delete(maps, "os")
	}

	if maps["probe_name"] != nil {
		slog.Println(slog.DEBUG, "probe_name", maps["probe_name"])
		dbTmp = dbTmp.Where("probe_result.probe_name LIKE ?", utils.GetInterfaceToString(maps["probe_name"])+"%")
		delete(maps, "probe_name")
	}

	if maps["task_name"] != nil {
		dbTmp = dbTmp.Where("task_id in ?", GetTaskIds(utils.GetInterfaceToString(maps["task_name"])))
		delete(maps, "task_name")
	}

	if maps["task_group"] != nil {
		dbTmp = dbTmp.Where("task_id in ?", GetTaskIdsBygrop(utils.GetInterfaceToString(maps["task_group"])))
		delete(maps, "task_group")
	}

	dbTmp.Where(maps).Count(&total)

	dbTmp.Where(maps).Offset(pageNum).Limit(pageSize).Order(order).Find(&ProbeRes)

	var ProbeResNew []define.ProbeRes

	pgrMap := GetPgRegionMap()

	pgMap := GetPgMap()

	osMap := GetOsMap(getipArr(ProbeRes))

	taskMap := GetTaskMap(gettaskArr(ProbeRes))

	// slog.Println(slog.DEBUG, "taskMap", taskMap)
	// probe_info.probe_send,probe_info.probe_recv,probe_info.probe_group,probe_info.probe_tags
	for _, v := range ProbeRes {

		// slog.Println(slog.DEBUG, "taskId", v.TaskID)
		v.Pg = pgMap[v.Pname].Group
		v.Payload = pgMap[v.Pname].Send
		v.Finger = pgMap[v.Pname].Recv
		v.Tags = pgMap[v.Pname].Tags
		v.Region = pgrMap[v.Pg]
		v.Os = osMap[v.IP]
		v.TaskName = taskMap[v.TaskID].Name
		v.TaskGroup = taskMap[v.TaskID].Group

		ProbeResNew = append(ProbeResNew, v)
	}

	return ProbeResNew, total
}

func getipArr(res []define.ProbeRes) (ipArr []string) {
	for _, v := range res {

		if utils.In_array(v.IP, ipArr) {
			continue
		}

		ipArr = append(ipArr, v.IP)
	}

	return
}

func gettaskArr(res []define.ProbeRes) (ipArr []string) {
	for _, v := range res {

		if utils.In_array(v.TaskID, ipArr) {
			continue
		}

		ipArr = append(ipArr, v.TaskID)
	}

	return
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
