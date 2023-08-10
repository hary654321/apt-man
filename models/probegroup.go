package models

import (
	"zrDispatch/core/utils/define"
)

// 创建任务
func AddProbeGroup(data define.ProbeGroupAdd) {

	ProbeGroupRes := define.ProbeGroupAdd{
		Type:   data.Type,
		Name:   data.Name,
		Region: data.Region,
		Desc:   data.Desc,
	}

	db.Table("probe_group").Create(&ProbeGroupRes)
}

func GetProbeGroup(pageNum int, pageSize int, maps map[string]interface{}) (probeGroup []define.ProbeGroupRes, total int64) {
	dbTmp := db.Table("probe_group")

	dbTmp.Where("is_deleted", 0).Where(maps).Offset(pageNum).Limit(pageSize).Order("probe_group_update_time  desc").Find(&probeGroup)

	dbTmp.Where("is_deleted", 0).Where(maps).Count(&total)
	return
}

func DeleteProbeGroup(ids []int) int64 {

	res := db.Table("probe_group").Where("probe_group_id in (?) ", ids).Update("is_deleted", 1)

	return res.RowsAffected
}

// 通过id，更新ip列表
func EditProbeGroupr(pge define.ProbeGroupE) int64 {

	res := db.Table("probe_group").Where("probe_group_id = ?", pge.ID).Updates(&pge)
	return res.RowsAffected
}

func GetGroupRegion() (pnr []define.ProbeGroupNr) {
	dbTmp := db.Table("probe_group")

	dbTmp.Select("probe_group_region", "probe_group_name").Find(&pnr)

	return
}

func GetPgSelect() (ProbeInfo []define.PgName) {
	dbTmp := db.Table("probe_group")

	dbTmp.Where("is_deleted", 0).Find(&ProbeInfo)

	return
}
