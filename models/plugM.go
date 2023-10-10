package models

import (
	"zrDispatch/core/slog"
	"zrDispatch/core/utils/define"
)

// 创建任务
func AddPlugInfo(pi define.PlugInfoAdd) error {

	res := db.Table("plug").Create(&pi)
	return res.Error
}

// 创建任务
func BatchAddPlugInfo(datas []map[string]any) int64 {
	res := db.Table("plug").Create(datas)
	return res.RowsAffected
}

func GetPlugInfo(pageNum int, pageSize int, maps map[string]interface{}) (PlugInfoRes []define.PlugInfo, total int64) {
	dbTmp := db.Table("plug")

	var PlugInfo []define.PlugInfo
	if maps["name"] != nil {
		dbTmp = dbTmp.Where("name LIKE ?", "%"+maps["name"].(string)+"%")
		delete(maps, "name")
	}

	dbTmp.Where("is_deleted", 0).Where(maps).Count(&total)
	dbTmp.Where("is_deleted", 0).Where(maps).Offset(pageNum).Limit(pageSize).Order("id  desc").Find(&PlugInfo)

	for _, v := range PlugInfo {
		v.StatusStr = v.Status.String()

		PlugInfoRes = append(PlugInfoRes, v)
	}

	return
}

func GetPlugInfoByName(name string, id string) (PlugMatch define.PlugInfoAdd) {
	dbTmp := db.Table("plug")

	dbTmp.Where("name = ?", name).Where("id != ?", id).Take(&PlugMatch)

	return
}

func GetPlugInfoById(id string) (PlugMatch define.PlugInfoAdd) {
	dbTmp := db.Table("plug")

	dbTmp.Where("id = ?", id).Take(&PlugMatch)

	return
}

func GetPlugSelect() (PlugInfo []define.PlugIdName) {
	dbTmp := db.Table("plug")

	dbTmp.Where("is_deleted", 0).Where("status", define.PLUG_SUCC).Order("update_time  desc").Find(&PlugInfo)

	return
}

func DeletePlugInfo(ids []int) int64 {

	res := db.Table("plug").Where("id in (?) ", ids).Update("is_deleted", 1)

	return res.RowsAffected
}

// 通过id，更新
func EditPlugInfo(pge define.PlugInfoE) error {

	res := db.Table("plug").Where("id = ?", pge.ID).Updates(pge)
	slog.Println(slog.DEBUG, res.Error)
	return res.Error
}

func ChangePlugState(name string, status int) error {

	res := db.Table("plug").Where("name", name).Update("status", status)

	return res.Error
}
