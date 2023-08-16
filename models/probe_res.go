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

func GetProbeRes(pageNum int, pageSize int, maps map[string]interface{}) (ProbeRes []define.ProbeRes, total int64) {
	dbTmp := db.Table("probe_result")

	dbTmp = dbTmp.Select("probe_result.*,probe_info.probe_send,probe_info.probe_recv,probe_info.probe_group,probe_info.probe_tags,probe_group.probe_group_region").
		Joins("left join probe_info on probe_info.probe_name = probe_result.probe_name").
		Joins("left join probe_group on probe_group.probe_group_name = probe_info.probe_group")

	if utils.GetInterfaceToString(maps["probe_group"]) != "" {
		dbTmp = dbTmp.Where("probe_info.probe_group = ?", utils.GetInterfaceToString(maps["probe_group"]))
		delete(maps, "probe_group")
	}

	if maps["probe_name"] != "" {
		dbTmp = dbTmp.Where("probe_result.probe_name LIKE ?", "%"+utils.GetInterfaceToString(maps["probe_name"])+"%")
		delete(maps, "probe_name")
	}

	dbTmp.Where(maps).Count(&total)

	dbTmp.Where(maps).Offset(pageNum).Limit(pageSize).Order("probe_result.id  desc").Find(&ProbeRes)

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

func GetTaskProbe(taskId string) (ProbeRes []define.ProbeRes) {
	dbTmp := db.Table("probe_result")

	dbTmp.Where("run_task_id like ? ", taskId+"%").Where("matched", define.Matched).Find(&ProbeRes)

	return
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
