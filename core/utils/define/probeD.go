package define

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type IdName struct {
	Id   string `gorm:"column:probe_id" json:"probe_id"`
	Name string `gorm:"column:probe_name" json:"probe_name" `
}

type ProbeInfoRes struct {
	Id int `gorm:"column:probe_id" json:"probe_id"`
	ProbeInfoAdd
	Ctime LocalTime `gorm:"column:probe_create_time" json:"probe_create_time"`
	Utime LocalTime `gorm:"column:probe_update_time" json:"probe_update_time"`
	D     int       `gorm:"column:is_deleted" json:"is_deleted"`
}

type ProbeGroupAdd struct {
	Name   string `gorm:"column:probe_group_name" json:"probe_group_name"`
	Type   string `gorm:"column:probe_group_type" json:"probe_group_type"`
	Region string `gorm:"column:probe_group_region" json:"probe_group_region"`
	Desc   string `gorm:"column:probe_group_desc" json:"probe_group_desc"`
}

type ProbeGroupE struct {
	ID int `gorm:"column:probe_group_id" json:"probe_group_id" binding:"required"`
	ProbeGroupAdd
}

type ProbeGroupNr struct {
	Name   string `gorm:"column:probe_group_name" json:"probe_group_name" `
	Region string `gorm:"column:probe_group_region" json:"probe_group_region" `
}

type ProbeGroupRes struct {
	Id int `gorm:"column:probe_group_id" json:"probe_group_id"`
	ProbeGroupAdd
	Ctime LocalTime `gorm:"column:probe_group_create_time" json:"probe_group_create_time"`
	Utime LocalTime `gorm:"column:probe_group_update_time" json:"probe_group_update_time"`
	D     int       `gorm:"column:is_deleted" json:"is_deleted"`
}

type PgName struct {
	Name string `gorm:"column:probe_group_name" json:"probe_group_name"`
}

type ProbeInfoAdd struct {
	Name  string `gorm:"column:probe_name" json:"probe_name" binding:"required"`
	Group string `gorm:"column:probe_group" json:"probe_group" binding:"required"`
	Tags  string `gorm:"column:probe_tags" json:"probe_tags" binding:"required"`
	Pro   string `gorm:"column:probe_protocol" json:"probe_protocol" binding:"required"`
	MT    string `gorm:"column:probe_match_type" json:"probe_match_type" binding:"required"`
	Send  string `gorm:"column:probe_send" json:"probe_send" binding:"required"`
	Recv  string `gorm:"column:probe_recv" json:"probe_recv" binding:"required"`
	Desc  string `gorm:"column:probe_desc" json:"probe_desc" binding:"required"`
}

type ProbeInfoE struct {
	ID int `gorm:"column:probe_id" json:"probe_id" binding:"required"`
	ProbeInfoAdd
}

type Pyload struct {
	Payload       string `gorm:"column:probe_send" json:"payload"`
	Name          string `gorm:"column:probe_name" json:"probe_name"`
	ProbeProtocol string `gorm:"column:probe_protocol" json:"probe_protocol"`
	Recv          string `gorm:"column:probe_recv" json:"probe_recv"`
	MT            string `gorm:"column:probe_match_type" json:"probe_match_type"`
}

// probe_name
type ProbeResCreate struct {
	Id        int    `gorm:"column:id" json:"id"`
	IP        string `gorm:"column:ip" json:"ip" `
	Pname     string `gorm:"column:probe_name" json:"probe_name" `
	Port      string `gorm:"column:port" json:"port" `
	Hex       string `gorm:"column:hex" json:"hex" `
	Res       string `gorm:"column:response" json:"response" `
	RunTaskID string `gorm:"column:run_task_id" json:"run_task_id"`
	Matched   int    `gorm:"column:matched" json:"matched"`
	Ctime     string `gorm:"column:create_time" json:"create_time"`
	D         int    `gorm:"column:is_deleted" json:"is_deleted"`

	// Utime   string `gorm:"column:update_time" json:"update_time"`
	// Type    string `gorm:"column:type" json:"type" `
}

type ProbeResEdit struct {
	Id     int    `gorm:"column:id" json:"id" binding:"required"`
	Dealed int    `gorm:"column:dealed" json:"dealed" `
	Remark string `gorm:"column:remark" json:"remark" `
	Utime  string `gorm:"column:update_time" json:"update_time"`
}

// probe_name
type ProbeRes struct {
	Id        int         `gorm:"column:id" json:"id"`
	IP        string      `gorm:"column:ip" json:"ip" `
	Pname     string      `gorm:"column:probe_name" json:"probe_name" `
	Pg        string      `gorm:"column:probe_group" json:"probe_group" `
	Tags      string      `gorm:"column:probe_tags" json:"probe_tags" `
	Region    string      `gorm:"column:probe_group_region" json:"probe_group_region" `
	Payload   string      `gorm:"column:probe_send" json:"payload" `
	Finger    string      `gorm:"column:probe_recv" json:"finger" `
	Port      string      `gorm:"column:port" json:"port" `
	Hex       string      `gorm:"column:hex" json:"hex" `
	Res       string      `gorm:"column:response" json:"response" `
	RunTaskID string      `gorm:"column:run_task_id" json:"run_task_id"`
	Matched   MatchStatus `gorm:"column:matched" json:"matched"`
	Dealed    DealStatus  `gorm:"column:dealed" json:"dealed"`
	Remark    string      `gorm:"column:remark" json:"remark"`
	Ctime     LocalTime   `gorm:"column:create_time" json:"create_time"`
	Utime     LocalTime   `gorm:"column:update_time" json:"update_time"`
	D         int         `gorm:"column:is_deleted" json:"is_deleted"`

	// Utime   string `gorm:"column:update_time" json:"update_time"`
	// Type    string `gorm:"column:type" json:"type" `
}

type MatchStatus int

const (
	MatchInit  MatchStatus = iota //初始化
	Matched                       // 匹配上
	NotMatched                    // 未匹配上
)

func (t MatchStatus) String() string {
	switch t {
	case MatchInit:
		return "初始化"
	case Matched:
		return "命中"
	case NotMatched:
		return "未命中"
	default:
		return "unknown"
	}
}

type DealStatus int

const (
	NotDeal DealStatus = iota //未处理
	DealEd                    // 已处理
)

func (t DealStatus) String() string {
	switch t {
	case NotDeal:
		return "未处理"
	case DealEd:
		return "已处理"
	default:
		return "unknown"
	}
}

type LocalTime time.Time

func (t LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	// LocalTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format("2006-01-02 15:04:05"), nil
}

func (t LocalTime) String() string {
	// LocalTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format("2006-01-02 15:04:05")
}

func (t *LocalTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = LocalTime(t1)
	return err
}
