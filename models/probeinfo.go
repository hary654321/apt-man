package models

import (
	"encoding/base64"
	"zrDispatch/core/slog"
	"zrDispatch/core/utils/define"
)

// 创建任务
func AddProbeInfo(pi define.ProbeInfoAdd) error {

	res := db.Table("probe_info").Create(&pi)
	return res.Error
}

// 创建任务
func BatchAddProbeInfo(datas []map[string]any) int64 {
	res := db.Table("probe_info").Create(datas)
	return res.RowsAffected
}

func GetProbeInfo(pageNum int, pageSize int, maps map[string]interface{}) (ProbeInfo []define.ProbeInfoRes, total int64) {
	dbTmp := db.Table("probe_info")

	dbTmp.Where("is_deleted", 0).Where(maps).Count(&total)
	dbTmp.Where("is_deleted", 0).Where(maps).Offset(pageNum).Limit(pageSize).Order("probe_id  desc").Find(&ProbeInfo)

	return
}

func GetProbeInfoByName(name string) (ProbeMatch define.ProbeInfoAdd) {
	dbTmp := db.Table("probe_info")

	dbTmp.Where("probe_name = ?", name).Take(&ProbeMatch)

	return
}

func GetProbeSelect() (ProbeInfo []define.IdName) {
	dbTmp := db.Table("probe_info")

	dbTmp.Where("is_deleted", 0).Order("probe_update_time  desc").Find(&ProbeInfo)

	return
}

func DeleteProbeInfo(ids []int) int64 {

	res := db.Table("probe_info").Where("probe_id in (?) ", ids).Update("is_deleted", 1)

	return res.RowsAffected
}

// 通过id，更新
func EditProbeInfo(pge define.ProbeInfoE) int64 {

	res := db.Table("probe_info").Where("probe_id = ?", pge.ID).Updates(pge)
	slog.Println(slog.DEBUG, res.Error)
	return res.RowsAffected
}

func GetPayload(ids []string) (p []define.Pyload, err error) {

	dbTmp := db.Table("probe_info")

	res := dbTmp.Where("probe_id in (?) ", ids).Order("probe_update_time  desc").Find(&p)

	for i := 0; i < len(p); i++ {
		p[i].Payload = base64.StdEncoding.EncodeToString([]byte(p[i].Payload))
	}
	err = res.Error
	return
}
