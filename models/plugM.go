package models

import (
	"encoding/base64"
	"zrDispatch/core/slog"
	"zrDispatch/core/utils/define"
)

// 创建任务
func AddPlugInfo(pi define.PlugInfoAdd) error {

	res := db.Table("Plug_info").Create(&pi)
	return res.Error
}

// 创建任务
func BatchAddPlugInfo(datas []map[string]any) int64 {
	res := db.Table("Plug_info").Create(datas)
	return res.RowsAffected
}

func GetPlugInfo(pageNum int, pageSize int, maps map[string]interface{}) (PlugInfo []define.PlugInfoRes, total int64) {
	dbTmp := db.Table("Plug_info")

	if maps["Plug_name"] != nil {
		dbTmp = dbTmp.Where("Plug_name LIKE ?", "%"+maps["Plug_name"].(string)+"%")
		delete(maps, "Plug_name")
	}

	dbTmp.Where("is_deleted", 0).Where(maps).Count(&total)
	dbTmp.Where("is_deleted", 0).Where(maps).Offset(pageNum).Limit(pageSize).Order("Plug_id  desc").Find(&PlugInfo)

	return
}

func GetPlugInfoByName(name string) (PlugMatch define.PlugInfoAdd) {
	dbTmp := db.Table("Plug_info")

	dbTmp.Where("Plug_name = ?", name).Take(&PlugMatch)

	return
}

func GetPlugSelect() (PlugInfo []define.IdName) {
	dbTmp := db.Table("Plug_info")

	dbTmp.Where("is_deleted", 0).Order("Plug_update_time  desc").Find(&PlugInfo)

	return
}

func DeletePlugInfo(ids []int) int64 {

	res := db.Table("Plug_info").Where("Plug_id in (?) ", ids).Update("is_deleted", 1)

	return res.RowsAffected
}

// 通过id，更新
func EditPlugInfo(pge define.PlugInfoE) int64 {

	res := db.Table("Plug_info").Where("Plug_id = ?", pge.ID).Updates(pge)
	slog.Println(slog.DEBUG, res.Error)
	return res.RowsAffected
}

func GetPayload(ids []string) (p []define.Pyload, err error) {

	dbTmp := db.Table("Plug_info")

	res := dbTmp.Where("Plug_id in (?) ", ids).Order("Plug_update_time  desc").Find(&p)

	for i := 0; i < len(p); i++ {
		p[i].Payload = base64.StdEncoding.EncodeToString([]byte(p[i].Payload))
	}
	err = res.Error
	return
}
