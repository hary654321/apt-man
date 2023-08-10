package models

import (
	"zrDispatch/core/utils/define"
)

// 创建任务
func AddCertRes(c define.Cert) int {

	res := db.Table("cert_result").Create(&c)
	return int(res.RowsAffected)
}

// 创建任务
func BatchAddCertRes(datas []map[string]any) int64 {

	// slog.Println(slog.DEBUG, datas)

	res := db.Table("cert_result").Create(datas)
	return res.RowsAffected
}

func GetCertRes(pageNum int, pageSize int, maps map[string]interface{}) (CertRes []define.CertRes, total int64) {
	dbTmp := db.Table("cert_result")

	dbTmp.Where(maps).Count(&total)

	dbTmp.Where(maps).Offset(pageNum).Limit(pageSize).
		Select("cert_result.*,probe_info.probe_recv").
		Joins("left join probe_info on probe_info.probe_name = cert_result.probe_name").Order("cert_result.id  desc").Find(&CertRes)

	return
}

func DeleteCertRes(ids []int) int64 {

	res := db.Table("cert_result").Where("id in (?) ", ids).Delete(&define.CertRes{})

	return res.RowsAffected
}
