package models

import (
	"zrDispatch/common/utils"
	"zrDispatch/core/utils/define"
)

func Gethost(hostname string) (host define.Host) {

	db.Table("host").Where("hostname = ? ", hostname).Take(&host)
	return
}

func UpdateHostStatus(id, status string) {

	db.Table("host").Where("id = ?", id).Update("status", status)
}

// 通过id，更新
func ChangeHost(host define.HostGorm) error {

	utils.WriteJsonLog(host)
	res := db.Table("host").Model(&define.HostGorm{}).Where("id = ?", host.ID).Updates(host)

	return res.Error
}
