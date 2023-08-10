package models

import (
	"zrDispatch/core/utils/define"
)

func AddMatchRes(c define.MatchRes) error {

	res := db.Table("match_result").Create(&c)
	return res.Error
}

func GetMatchRes(pageNum int, pageSize int, maps map[string]interface{}) (res []define.MatchRes, total int64) {
	dbTmp := db.Table("match_result")

	dbTmp.Where(maps).Count(&total)

	dbTmp.Where(maps).Offset(pageNum).Limit(pageSize).Order("id  desc").Find(&res)

	return
}
