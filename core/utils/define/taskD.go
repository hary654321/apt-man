package define

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

// DataCode run code
type TaskData struct {
	T   string `json:"t"`
	SPY string `json:"spy"`
}

// AlarmStatus task is alarm
type AlarmStatus int8

const (
	// All will alarm after task run
	All AlarmStatus = -2
	// Fail will alarm after task fail
	Fail AlarmStatus = -1
	// Success will alarm after task success
	Success AlarmStatus = 1
)

func (al AlarmStatus) String() string {
	switch al {
	case All:
		return "All"
	case Fail:
		return "Fail"
	case Success:
		return "Success"
	default:
		return "Unknown"
	}
}

// CreateTask struct
type CreateTask struct {
	GetID
	GetName
	Task
}

// CreateTask struct
type DetailTask struct {
	GetID
	GetName
	Task
	RID
	Time
	Status TaskOneStatus `json:"status" comment:"运行状态码"`
}

type Time struct {
	Ctime LocalTime `gorm:"column:createTime" json:"create_time"`
	Utime LocalTime `gorm:"column:updateTime" json:"update_time"`
}

// ChangeTask struct
type ChangeTask struct {
	IDName
	Task
}

// IDName struct
type IDName struct {
	GetID
	GetName
}

// TaskOneStatus=====================ssssss============
type TaskOneStatus int

const (
	// Email check email
	TASK_STATUS_INIT TaskOneStatus = iota
	// Name check name
	TASK_STATUS_RUNING
	//
	TASK_STATUS_DONE

	TASK_STATUS_Fail

	TASK_STATUS_STOP

	TASK_STATUS_GQ

	TASK_STATUS_WAIT
)

func (s TaskOneStatus) String() string {
	switch s {
	case TASK_STATUS_INIT:
		return "初始化"
	case TASK_STATUS_RUNING:
		return "执行中"
	case TASK_STATUS_DONE:
		return "执行完成"
	case TASK_STATUS_Fail:
		return "执行失败"
	case TASK_STATUS_STOP:
		return "终止执行"
	case TASK_STATUS_GQ:
		return "挂起"
	case TASK_STATUS_WAIT:
		return "待执行"
	default:
		return "unknow"
	}
}

// TaskOneStatus=====================eeeeee============

// GetTask get task
type GetTask struct {
	Group             string        `json:"group" comment:"分组"`
	TaskType          TaskType      `json:"task_type"`
	TaskTypeDesc      string        `json:"task_typedesc" comment:"任务类型"`
	Ip                string        `json:"ip" comment:"任务数据"`
	Port              string        `json:"port" comment:"任务数据"`
	Run               bool          `json:"run" comment:"运行"`
	Status            TaskOneStatus `json:"status" comment:"运行状态码"`
	StatusDesc        string        `json:"status_desc" comment:"运行状态"`
	ParentTaskIds     StrArr        `json:"parent_taskids"`
	ParentTaskIdsDesc StrArr        `json:"parent_taskidsdesc" comment:"父任务"`
	ParentRunParallel bool          `json:"parent_runparallel" comment:"父任务运行策略"`
	ChildTaskIds      StrArr        `json:"child_taskids"`
	ChildTaskIdsDesc  StrArr        `json:"child_taskidsdesc"  comment:"子任务"`
	ChildRunParallel  bool          `json:"child_runparallel" comment:"子任务运行策略"`
	CreateBy          string        `json:"create_by"`
	CreateByUID       string        `json:"create_byuid"`
	HostGroup         string        `json:"host_group" comment:"主机组"`
	HostGroupID       string        `json:"host_groupid"`
	Priority          int           `json:"priority" comment:"优先级"`
	Cronexpr          string        `json:"cronexpr" comment:"CronExpr"`
	Timeout           int           `json:"timeout" comment:"超时时间"`
	Threads           int           `json:"threads" comment:"超时时间"`
	AlarmUserIds      StrArr        `json:"alarm_userids"`
	AlarmUserIdsDesc  StrArr        `json:"alarm_useridsdesc" comment:"报警用户"`
	RoutePolicy       RoutePolicy   `json:"route_policy"`
	RoutePolicyDesc   string        `json:"route_policydesc" comment:"路由策略"`
	ExpectCode        int           `json:"expect_code"  comment:"期望返回码"`
	ExpectContent     string        `json:"expect_content" comment:"期望返回内容"`
	AlarmStatus       AlarmStatus   `json:"alarm_status"`
	AlarmStatusDesc   string        `json:"alarm_statusdesc" comment:"报警策略"`
	ProbeId           StrArr        `json:"probeId" comment:"规则"`
	Plug              StrArr        `json:"plug" comment:"插件"`
	Common
}

type StrArr []string

//Implement Scanner interface

func (t *StrArr) Scan(value interface{}) error {

	// slog.Println(slog.DEBUG, value)

	if value == nil {
		*t = StrArr{}
	}

	val, ok := value.([]byte)

	if !ok {

		return errors.New(fmt.Sprint("wrong type", value))

	}

	*t = StrArr(Split(string(val)))

	return nil

}

func Split(a string) (res []string) {
	arr := strings.Split(a, ",")
	for _, v := range arr {
		if v != "" {
			res = append(res, v)
		}
	}
	return
}

//Implement Valuer interface

func (t StrArr) Value() (driver.Value, error) {

	//this check is here if you don't want to save an empty string

	if len(t) == 0 {

		return "", nil

	}

	return []byte(strings.Join(t, ",")), nil

}

// RoutePolicy set a task hot to select run worker
type RoutePolicy uint8

const (
	// Random get host by random
	Random RoutePolicy = iota + 1
	// RoundRobin get host by order
	RoundRobin
	// Weight get host by host weight
	Weight
	// LeastTask get host by host LeastTask
	LeastTask
)

func (r RoutePolicy) String() string {
	switch r {
	case Random:
		return "随机"
	case RoundRobin:
		return "轮询"
	case Weight:
		return "权重"
	case LeastTask:
		return "LeastTask"
	default:
		return "未知"
	}
}

// Trigger return how to trigger run task
type Trigger uint8

const (
	// Auto cron run task
	Auto Trigger = iota + 1
	// Manual trigger run task
	Manual
)

func (t Trigger) String() string {
	switch t {
	case Auto:
		return "自动触发"
	case Manual:
		return "手动触发"
	default:
		return "UnKnown"
	}
}

// RunTask running task message
type RunTask struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Cronexpr     string  `json:"cronexpr"`
	StartTimeStr string  `json:"start_timestr"`
	StartTime    int64   `json:"start_time"` // use ms,
	RunTime      int     `json:"run_time"`   // s
	Trigger      Trigger `json:"trigger"`
	TriggerStr   string  `json:"triggerstr"`
}

// TaskResp run task resp message
type TaskResp struct {
	TaskID      string       `json:"task_id"`
	Task        string       `json:"task"`
	LogData     string       `json:"resp_data"`    // task run log data
	Code        int          `json:"code"`         // return code
	TaskType    TaskRespType `json:"task_type"`    // 1 主任务 2 父任务 3 子任务
	TaskTypeStr string       `json:"task_typestr"` // 1 主任务 2 父任务 3 子任务
	RunHost     string       `json:"run_host"`     // task run host
	Status      string       `json:"status"`       // task status finish,fail, cancel
}

// Task define Task

type Task struct {
	Group             string   `gorm:"column:group"  json:"group" form:"group"  binding:"required,min=1,max=30"`
	TaskType          TaskType `gorm:"column:taskType"  json:"task_type"  binding:"required" ` // 任务类型
	Ip                string   `gorm:"column:ip"  json:"ip" binding:"required"  `              // 任务数据
	Port              string   `gorm:"column:port"  json:"port" `
	Run               int      `gorm:"column:run" json:"run" `                             // 是否可以自动调度  如果为false则只能手动或者被其他任务依赖运行
	ParentRunParallel bool     `gorm:"column:parentRunParallel" json:"parent_runparallel"` // 是否以并行运行父任务 0否 1是
	ChildRunParallel  bool     `gorm:"column:childRunParallel" json:"child_runparallel"`   // 是否以并行运行子任务 否 1是
	// CreateBy          string      `gorm:"column:create_by" json:"create_by"`                                 // 创建人
	CreateByUID string `gorm:"column:createByID" json:"create_byuid"` // 创建人ID
	// HostGroup    string      `gorm:"column:host_group" json:"host_group" `                              // 主机组
	HostGroupID string      `gorm:"column:hostGroupID" json:"host_groupid" binding:"required"` // 主机组ID
	Cronexpr    string      `gorm:"column:cronExpr" json:"cronexpr" binding:"max=1000"`        // 执行任务表达式
	Priority    int         `gorm:"column:priority" json:"priority" binding:"required,min=-1"`
	Timeout     int         `gorm:"column:timeout" json:"timeout" binding:"required,min=-1"`        // 任务超时时间 (s) -1 no limit
	RoutePolicy RoutePolicy `gorm:"column:routePolicy" json:"route_policy" "required,min=1,max=4" ` // how to select a run worker from hostgroup
	//ExpectCode        int         `json:"expect_code"`                                  // expect task return code. if not set 0 or 200
	//ExpectContent     string      `json:"expect_content"`                               // expect task return content. if not set do not check
	AlarmStatus AlarmStatus `gorm:"column:alarmStatus" json:"alarm_status" binding:"required,min=-2,max=1"` // alarm when task run success or fail or all all:-2 failed: -1 success: 1
	Remark      string      `gorm:"column:remark" json:"remark" binding:"max=100"`
	ProbeId     StrArr      `gorm:"column:probeId" json:"probeId" comment:"规则"`
	Plug        StrArr      `gorm:"column:plug" json:"plug" comment:"插件"`
	Threads     int         `gorm:"column:threads" json:"threads" binding:"required,min=-1"` // 任务超时时间 (s) -1 no limit
}

type RID struct {
	RunTaskId string `gorm:"column:run_task_id" json:"run_task_id"`
}

type TaskArrEle struct {
	ChildTaskIds  []string `gorm:"column:childTaskIds" json:"child_taskids" binding:"max=20"` // 子任务 运行结束后运行子任务
	ParentTaskIds []string `gorm:"column:cronExpr" json:"parent_taskids" binding:"max=20"`    // 父任务 运行任务前先运行父任务 以父或子任务运行时 任务不会执行自已的父子任务，防止循环依赖
	AlarmUserIds  []string `gorm:"column:alarmUserIds" json:"alarm_userids" `                 // 报警用户 最多十个多个用户
}

// Task define Task
type TaskE struct {
	Id string `json:"id" binding:"id"` // 任务数据
	Task
}

// TaskType task type
// shell
// api
// TaskDataType
type TaskType uint8

const (
	// Code run code
	TYPE_PORT TaskType = iota + 1
	// API run http req
	TYPE_PROBE
)

func (tt TaskType) String() string {
	switch tt {
	case TYPE_PORT:
		return "端口扫描"
	case TYPE_PROBE:
		return "规则扫描"
	default:
		return "unknow"
	}
}

func (tt TaskType) Value() string {
	switch tt {
	case TYPE_PORT:
		return "port"
	case TYPE_PROBE:
		return "probe"
	default:
		return "unknow"
	}
}
