package models

import (
	"zrDispatch/common/utils"
	"zrDispatch/core/utils/define"
)

// 创建任务
func AddProbeGroup(data define.ProbeGroupAdd) {

	ProbeGroupRes := define.ProbeGroupAdd{
		Type:   data.Type,
		Name:   data.Name,
		Region: data.Region,
		Desc:   data.Desc,
		Ctime:  utils.GetTimeStr(),
	}

	db.Table("probe_group").Create(&ProbeGroupRes)
}

func GetProbeGroup(pageNum int, pageSize int, maps map[string]interface{}) (probeGroup []define.ProbeGroupRes, total int64) {
	dbTmp := db.Table("probe_group")

	dbTmp.Where("is_deleted", 0).Where(maps).Count(&total)

	dbTmp.Where("is_deleted", 0).Where(maps).Offset(pageNum).Limit(pageSize).Order("probe_group_id  desc").Find(&probeGroup)

	return
}

func GetProbeGroupByID(id int) (probeGroup *define.ProbeGroupRes) {
	dbTmp := db.Table("probe_group")

	dbTmp.Where("probe_group_id = ?", id).Where("is_deleted", 0).Take(&probeGroup)

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

func GetPgRegionMap() map[string]string {
	dbTmp := db.Table("probe_group")

	var ProbeInfo []define.PGR

	dbTmp.Find(&ProbeInfo)

	pgmap := make(map[string]string)
	for _, v := range ProbeInfo {
		pgmap[v.Name] = v.Region
	}

	return pgmap
}
