package models

import (
	"zrDispatch/core/utils/define"
)

func AddPlugRes(c define.PlugResAdd) error {

	res := db.Table("plug_result").Create(&c)
	return res.Error
}

func GetPlugRes(pageNum int, pageSize int, maps map[string]interface{}) (res []define.PlugRes, total int64) {
	dbTmp := db.Table("plug_result")

	dbTmp.Where(maps).Count(&total)

	dbTmp.Where(maps).Offset(pageNum).Limit(pageSize).Order("id  desc").Find(&res).Scan(&res)

	for i := 0; i < len(res); i++ {
		res[i].TypeDesc = res[i].Type.String()
	}

	return
}
