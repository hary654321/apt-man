package define

type PlugResAdd struct {
	RunTaskID string   `gorm:"column:run_task_id" json:"run_task_id"`
	Type      TaskType `gorm:"column:type" json:"type" `
	Res       string   `gorm:"column:res" json:"res" `
	Ctime     string   `gorm:"column:create_time" json:"create_time"`
}

type PlugRes struct {
	GetID
	PlugResAdd
	Ctime    LocalTime `gorm:"column:create_time" json:"create_time"`
	Utime    LocalTime `gorm:"column:update_time" json:"update_time"`
	TypeDesc string    `json:"type_desc" `
}
