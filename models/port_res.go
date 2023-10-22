package models

import (
	"zrDispatch/core/slog"
	"zrDispatch/core/utils/define"
)

// 保存结果
func AddPortRes(c define.PortScan) int {

	if c.Os != "" {
		AddOs(Os{IP: c.IP, Os: c.Os})
	}

	res := db.Table("port_result").Create(&c)
	return int(res.RowsAffected)
}

func GetPortRes(pageNum int, pageSize int, maps map[string]interface{}) (PortRes []define.PortRes, total int64) {
	dbTmp := db.Table("port_result")

	dbTmp.Where(maps).Count(&total)

	dbTmp.Where(maps).Offset(pageNum).Limit(pageSize).Order("id  desc").Find(&PortRes)

	return
}

func DeletePortRes(ids []int) int64 {

	res := db.Table("port_result").Where("probe_id in (?) ", ids).Delete(&define.PortRes{})

	return res.RowsAffected
}

func GetTaskPortRes(taskId string) []define.PortResJJ {
	dbTmp := db.Table("port_result")

	var PortRes []define.PortResJJ
	dbTmp.Select("ip,port,os,service,product_name,probe_name,version").Where("run_task_id like ? ", taskId+"%").Order("id  desc").Find(&PortRes)

	havemap := make(map[string]int)

	var PortResUnique []define.PortResJJ

	for _, v := range PortRes {

		if havemap[v.IP+v.Port] == 1 {
			continue
		}

		PortResUnique = append(PortResUnique, v)

		havemap[v.IP+v.Port] = 1
	}

	return PortResUnique
}

type Result struct {
	Service string
	Total   int
}

func GetTaskPortGroup(taskId string) (PortRes []Result, live_port int) {
	dbTmp := db.Table("port_result")

	dbTmp.Select("count(DISTINCT(`ip`))  as total ,service").Group("service").Where("task_id = ? ", taskId).Scan(&PortRes)

	for _, v := range PortRes {
		live_port += v.Total
	}
	return
}

func GetTaskLiveIpCount(taskId string) (ipcount int64) {
	dbTmp := db.Table("port_result")

	dbTmp.Where("run_task_id like ? ", taskId+"%").Select("distinct ip").Distinct("ip").Count(&ipcount)

	return
}

func GetOsSelect() (os []define.Os) {
	dbTmp := db.Table("port_result")

	dbTmp.Where("os!=?", "").Select("os").Distinct("os").Find(&os)

	return
}

// 通过id，更新
func EditPortRes(pge define.PortResEdit) int64 {

	res := db.Table("port_result").Where("id = ?", pge.Id).Updates(pge)
	slog.Println(slog.DEBUG, res.Error)
	return res.RowsAffected
}
