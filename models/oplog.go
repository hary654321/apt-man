package models

import (
	"zrDispatch/core/utils/define"
)

func GetOpLogs(pageNum int, pageSize int, maps map[string]interface{}) (PortRes []define.OperateLog, total int64) {
	dbTmp := db.Table("port_result")

	dbTmp.Where(maps).Count(&total)

	dbTmp.Where(maps).Offset(pageNum).Limit(pageSize).Order("id  desc").Find(&PortRes)

	return
}
