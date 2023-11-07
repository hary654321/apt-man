package models

import (
	"zrDispatch/core/slog"
	"zrDispatch/core/utils/define"
)

// 创建任务
func AddOs(pi define.OsAdd) error {

	res := db.Table("os").Create(&pi)
	return res.Error
}

func IPCont(ip string) int64 {
	dbTmp := db.Table("os")
	var os define.Os
	res := dbTmp.Where("ip=?", ip).First(&os)
	slog.Println(slog.DEBUG, "res.RowsAffected", res.RowsAffected)
	return res.RowsAffected
}

func GetOsSelect() (os []define.Os) {
	dbTmp := db.Table("os")

	dbTmp.Where("os!=?", "").Select("os").Distinct("os").Find(&os)

	return
}

func GetIpRes(pageNum int, pageSize int, maps map[string]interface{}) (PortRes []define.OsRes, total int64) {
	dbTmp := db.Table("os")

	dbTmp.Where(maps).Count(&total)

	dbTmp.Where(maps).Offset(pageNum).Limit(pageSize).Order("id  desc").Find(&PortRes)

	return
}
