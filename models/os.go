package models

type Os struct {
	IP string `gorm:"column:ip" json:"ip"`
	Os string `gorm:"column:os" json:"os"`
}

// 创建任务
func AddOs(pi Os) error {

	res := db.Table("os").Create(&pi)
	return res.Error
}
