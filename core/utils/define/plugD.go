package define

type PlugResAdd struct {
	RunTaskID string `gorm:"column:run_task_id" json:"run_task_id"`
	Plug      string `gorm:"column:plug" json:"plug" `
	Res       string `gorm:"column:res" json:"res" `
	Ctime     string `gorm:"column:create_time" json:"create_time"`
}

type PlugRes struct {
	GetID
	PlugResAdd
	Ctime    LocalTime `gorm:"column:create_time" json:"create_time"`
	Utime    LocalTime `gorm:"column:update_time" json:"update_time"`
	TypeDesc string    `json:"type_desc" `
}

type PlugInfoAdd struct {
	Name     string     `gorm:"column:name"  json:"name" form:"name"  binding:"required,min=1,max=30"`
	FileName string     `gorm:"column:filename"  json:"filename" form:"filename"`
	Cmd      string     `gorm:"column:cmd"  json:"cmd" form:"cmd"  binding:"required,min=1,max=255"`
	Desc     string     `gorm:"column:desc"  json:"desc" form:"desc"  binding:"required,min=1,max=300"`
	Ctime    string     `gorm:"column:create_time" json:"create_time"`
	Sys      int        `gorm:"column:sys" json:"sys"`
	Status   PlugStatus `gorm:"column:status" json:"status"`
}

type PlugInfo struct {
	GetID
	PlugInfoAdd
	StatusStr string    `gorm:"column:status_text" json:"status_text"`
	Ctime     LocalTime `gorm:"column:create_time" json:"create_time"`
	Utime     LocalTime `gorm:"column:update_time" json:"update_time"`
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

type PlugStatus int

const (
	PLUG_JYZ PlugStatus = iota + 1
	PLUG_FAIL
	PLUG_SUCC
)

func (t PlugStatus) String() string {
	switch t {
	case PLUG_JYZ:
		return "校验中"
	case PLUG_FAIL:
		return "校验失败"
	case PLUG_SUCC:
		return "校验成功"
	default:
		return "unknown"
	}
}
