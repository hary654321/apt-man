package models

import (
	"encoding/base64"
	"zrDispatch/common/utils"
	"zrDispatch/core/slog"
	"zrDispatch/core/utils/define"
)

// 创建任务
func AddProbeInfo(pi define.ProbeInfoAdd) error {

	pi.Ctime = utils.GetTimeStr()
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

	if maps["probe_name"] != nil {
		dbTmp = dbTmp.Where("probe_name LIKE ?", "%"+maps["probe_name"].(string)+"%")
		delete(maps, "probe_name")
	}

	dbTmp.Where("is_deleted", 0).Where(maps).Count(&total)
	dbTmp.Where("is_deleted", 0).Where(maps).Offset(pageNum).Limit(pageSize).Order("probe_id  desc").Find(&ProbeInfo)

	return
}

func GetProbeInfoByName(name string) (ProbeMatch define.ProbeInfoAdd) {
	dbTmp := db.Table("probe_info")

	dbTmp.Where("probe_name = ?", name).Take(&ProbeMatch)

	return
}

func GetProbeInfoByPName(pgname string) (ProbeMatch define.ProbeInfoAdd) {
	dbTmp := db.Table("probe_info")

	dbTmp.Where("probe_group = ?", pgname).Take(&ProbeMatch)

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
		// p[i].PortArr = utils.GetPortArr(p[i].Port)
	}
	err = res.Error
	return
}

func GetPgMap() map[string]define.ProbeInfoAdd {
	dbTmp := db.Table("probe_info")

	var ProbeInfo []define.ProbeInfoAdd

	dbTmp.Find(&ProbeInfo)

	pgmap := make(map[string]define.ProbeInfoAdd)
	for _, v := range ProbeInfo {
		pgmap[v.Name] = v
	}

	return pgmap
}

func GetPname(pg string) (res []string) {
	dbTmp := db.Table("probe_info")

	var ProbeInfo []define.ProbeName
	dbTmp.Where("probe_group", pg).Find(&ProbeInfo)

	for _, v := range ProbeInfo {
		res = append(res, v.Name)
	}
	return
}
