package models

import (
	"zrDispatch/common/utils"
	"zrDispatch/core/utils/define"
)

func AddPlugRes(c define.PlugResAdd) error {

	res := db.Table("plug_result").Create(&c)
	return res.Error
}

func GetTaskPlugRes(taskId string) []define.PlugRes {
	dbTmp := db.Table("plug_result")

	var PlugRes []define.PlugRes
	dbTmp.Where("task_id = ? ", taskId).Order("id  desc").Find(&PlugRes)

	havemap := make(map[string]int)

	var PlugResUnique []define.PlugRes

	for _, v := range PlugRes {

		runid := v.RunTaskID
		tid := utils.SubString(runid, "", "-") + v.Plug
		if havemap[tid] == 1 {
			continue
		}

		PlugResUnique = append(PlugResUnique, v)

		havemap[tid] = 1
	}

	return PlugResUnique
}
