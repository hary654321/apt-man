package define

type OsAdd struct {
	IP    string `gorm:"column:ip" json:"ip"`
	Os    string `gorm:"column:os" json:"os"`
	Port  string `gorm:"column:port" json:"port"`
	Ctime string `gorm:"column:create_time" json:"create_time"`
}

type OsRes struct {
	Id    string    `gorm:"column:id" json:"id"`
	IP    string    `gorm:"column:ip" json:"ip"`
	Os    string    `gorm:"column:os" json:"os"`
	Port  string    `gorm:"column:port" json:"port"`
	Ctime LocalTime `gorm:"column:create_time" json:"create_time"`
}
