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

type PlugInfoAdd struct {
	Name     string `gorm:"column:name"  json:"name" form:"name"  binding:"required,min=1,max=30"`
	FileName string `gorm:"column:filename"  json:"filename" form:"filename"`
	Cmd      string `gorm:"column:cmd"  json:"cmd" form:"cmd"  binding:"required,min=1,max=255"`
	Desc     string `gorm:"column:desc"  json:"desc" form:"desc"  binding:"required,min=1,max=300"`
	Sys      int    `gorm:"column:sys" json:"sys"`
}

type PlugInfo struct {
	GetID
	PlugInfoAdd
	Ctime LocalTime `gorm:"column:create_time" json:"create_time"`
	Utime LocalTime `gorm:"column:update_time" json:"update_time"`
}

type PlugInfoE struct {
	ID       string `gorm:"column:id"  json:"id" form:"id" `
	Name     string `gorm:"column:name"  json:"name" form:"name"  binding:"required,min=1,max=30"`
	FileName string `gorm:"column:filename"  json:"filename" form:"filename"`
	Cmd      string `gorm:"column:cmd"  json:"cmd" form:"cmd"  binding:"required,min=1,max=255"`
	Desc     string `gorm:"column:desc"  json:"desc" form:"desc"  binding:"required,min=1,max=300"`
}

type PlugIdName struct {
	Id   string `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name" `
}
