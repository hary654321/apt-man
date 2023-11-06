package models

import "zrDispatch/core/utils/define"

// 创建任务
func AddOs(pi define.OsAdd) error {

	res := db.Table("os").Create(&pi)
	return res.Error
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
