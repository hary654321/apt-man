package define

const (
	// DefaultLimit set get total page
	DefaultLimit = 20
	// DefaultLimit set get total page
	ExportLimit = 50000
)

// Role Admin or Normal User
type Role uint8

const (
	// NormalUser define normal user
	NormalUser Role = iota + 1 // 普通用户 只对自已创建的主机或者主机组具有操作权限
	// AdminUser define admin user
	AdminUser // 管理员 具有所有操作
	// GuestUser only look
	GuestUser // 访客 只有查看的权限
)

// 这里不能瞎改的  和权限表对应的
func (r Role) String() string {
	switch r {
	case AdminUser:
		return "Admin"
	case NormalUser:
		return "Normal"
	case GuestUser:
		return "Guest"
	default:
		return "Unknown"
	}
}

// RunMode crocodile run mode
// run crocodile as server or client
type RunMode uint8

const (
	// Server run crocodile as server
	Server RunMode = iota + 1
	// Client run crocodile as client
	Client
)

// TaskRespType task type (parent task,master task, child task)
// TODO Rename TaskRunType
type TaskRespType uint8

const (
	// MasterTask task as master run
	MasterTask TaskRespType = iota + 1
	// ParentTask task as a task's parent task run
	ParentTask
	// ChildTask task as a task's child task run
	ChildTask
)

func (tasktype TaskRespType) String() string {
	switch tasktype {
	case MasterTask:
		return "master"
	case ChildTask:
		return "child"
	case ParentTask:
		return "parent"
	default:
		return "unknown"
	}
}

// GetID get task id in post
type GetID struct {
	ID string `gorm:"column:id"  json:"id" form:"id" `
}

// GetID get task id in post
type GetIdChange struct {
	ID  string `gorm:"column:id"  json:"id" form:"id" `
	RUN int    `gorm:"column:run"  json:"run" form:"run" `
}

type GetIDInt struct {
	ID int `gorm:"column:id"  json:"id" form:"id" `
}

// GetName get task name in post
type GetName struct {
	Name string `gorm:"column:name"  json:"name" form:"name"  binding:"required,min=1,max=30"`
}

// Common struct
type Common struct {
	ID         string `json:"id" comment:"ID"`
	Name       string `json:"name,omitempty" comment:"名称"`
	CreateTime string `json:"create_time,omitempty" comment:"创建时间"` // 创建时间
	UpdateTime string `json:"update_time,omitempty" comment:"更新时间"` // 最后一次更新时间
	Remark     string `json:"remark" comment:"备注"`                  // 备注
}

// User Struct
type User struct {
	Role      Role     `json:"role"`                               // 用户类型: 1 普通用户 2 管理员 3访客
	Roles     []string `json:"roles"`                              // 管理员
	RoleStr   string   `json:"rolestr,omitempty" comment:"用户类型"`   // 用户类型
	Forbid    bool     `json:"forbid" comment:"禁止用户"`              // 禁止用户登陆
	Password  string   `json:"password,omitempty" comment:"密码"`    // 用户密码
	Email     string   `json:"email" binding:"email" comment:"邮箱"` // 用户邮箱 日后任务的通知信息会发送给此邮件
	WeChat    string   `json:"wechat" comment:"WeChat"`            // wechat id
	DingPhone string   `json:"dingphone" comment:"钉钉"`             // dingding phone
	Slack     string   `json:"slack" comment:"Slack"`              // slack user name
	Telegram  string   `json:"telegram" comment:"Telegram"`        // telegram bot chat id
	Common
}

// RegistryUser data
type RegistryUser struct {
	Name     string `json:"name" binding:"required,max=30"`      // 用户名
	Password string `json:"password" binding:"required,min=8"`   // 用户密码
	Role     Role   `json:"role" binding:"required,min=1,max=3"` // 用户类型: 1 普通用户 2 管理员
	Remark   string `json:"remark" binding:"max=100"`            // 备注
}

// CreateAdminUser first run must be create admin user
type CreateAdminUser struct {
	Name     string `json:"username" binding:"required,max=30"` // 用户名
	Password string `json:"password" binding:"required,min=8"`  // 用户密码
}

// AdminChangeUser struct
type AdminChangeUser struct {
	ID       string `json:"id"  binding:"required,len=18"`       // user id
	Role     Role   `json:"role" binding:"required,min=1,max=3"` // 用户类型: 1 普通用户 2 管理员
	Forbid   bool   `json:"forbid"`                              // 禁止用户: 1 未禁止 2 禁止登陆
	Password string `json:"password"`                            // 用户密码 Common
	Remark   string `json:"remark"`                              // 备注 Common
}

// ChangeUserSelf change self's config
type ChangeUserSelf struct {
	ID        string `json:"id"  binding:"required"`  // user id
	Name      string `json:"name" binding:"required"` // 用户名称
	Email     string `json:"email"`                   // 用户邮箱
	WeChat    string `json:"wechat"`                  // wechat id
	DingPhone string `json:"dingphone"`               // dingding phone
	Telegram  string `json:"telegram"`                // telegram bot chat id
	Password  string `json:"password"`
	Remark    string `json:"remark"`
}

// Log task log
type Log struct {
	Name           string       `json:"name"`                // task log
	TaskID         string       `json:"taskid"`              // run taskid
	RunTaskID      string       `json:"runTaskId"`           // run taskid
	HostId         string       `json:"hostid"`              // run taskid
	StartTime      string       `json:"start_time"`          // ms
	StartTimeStr   string       `json:"start_timestr"`       //
	EndTime        string       `json:"end_time"`            // ms
	EndTimeStr     string       `json:"end_timestr"`         //
	TotalRunTime   int          `json:"total_runtime"`       // ms
	Status         int          `json:"status"`              // 任务运行结果 -1 失败 0进行中 1 成功
	Progress       int          `json:"progress"`            // 任务进度
	Taskresps      string       `json:"taskresps,omitempty"` // 任务执行过程日志
	Trigger        Trigger      `json:"trigger"`             // 任务触发
	Triggerstr     string       `json:"trigger_str"`         // 任务触发
	ErrCode        int          `json:"err_code"`            // err code
	ErrMsg         string       `json:"err_msg"`             // 错误原因
	ErrTasktype    TaskRespType `json:"err_tasktype"`        // err task type
	ErrTaskTypeStr string       `json:"err_tasktypestr"`     // 1 主任务 2 父任务 3 子任务
	ErrTaskID      string       `json:"err_taskid"`          // task failed id
	ErrTask        string       `json:"err_task"`            // task failed id
}

// Cleanlog data
type Cleanlog struct {
	GetName
	PreDay int64 `json:"preday"` // preday几天前的日志 0 为全部日志
}

// Query recv url query params
type Query struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}

// KlOption vue el-select
type KlOption struct {
	Label  string `json:"label"`
	Value  string `json:"value"`
	Online int    `json:"online,omitempty"` // online: 1 offline: -1
}

// TaskStatusTree real task tree
type TaskStatusTree struct {
	Name         string       `json:"name"`
	ID           string       `json:"id,omitempty"`
	Status       string       `json:"status"`
	TaskType     TaskRespType `json:"tasktype"`
	TaskRespData string       `json:"taskresp_data,omitempty"`
	// RunHost      string            `json:"runhost,omitempty"`
	Children []*TaskStatusTree `json:"children,omitempty"`
}

// GetTasksTreeStatus return a slice
func GetTasksTreeStatus() []*TaskStatusTree {
	retTasksStatus := make([]*TaskStatusTree, 0, 3)
	parentTasksStatus := &TaskStatusTree{
		Name:     "ParentTasks",
		Status:   TsNoData.String(),
		Children: make([]*TaskStatusTree, 0),
	}

	mainTaskStatus := &TaskStatusTree{
		// Name:   task.name,
		// ID:     taskid,
		TaskType: MasterTask,
		Status:   TsNoData.String(),
	}

	childTasksStatus := &TaskStatusTree{
		Name:     "ChildTasks",
		Status:   TsNoData.String(),
		Children: make([]*TaskStatusTree, 0),
	}

	retTasksStatus = append(retTasksStatus,
		parentTasksStatus,
		mainTaskStatus,
		childTasksStatus)
	return retTasksStatus
}

// TaskStatus task run status
type TaskStatus uint

const (
	// TsWait task is waiting pre task is running
	TsWait TaskStatus = iota + 1
	// TsRun tassk is running
	TsRun
	// TsFinish task is run finish
	TsFinish
	// TsFail task run fail
	TsFail
	// TsCancel task is cancel ,because pre task is run fail
	TsCancel
	// TsNoData parenttasks or childtasks no task
	TsNoData
)

func (t TaskStatus) String() string {
	switch t {
	case TsWait:
		return "wait"
	case TsRun:
		return "run"
	case TsFinish:
		return "finish"
	case TsFail:
		return "fail"
	case TsCancel:
		return "cancel"
	case TsNoData:
		return "nodata"
	default:
		return "unknown"
	}
}

// OperateLog openrate log
type OperateLog struct {
	UID         string   `json:"user_id"`      // 修改人ID
	UserName    string   `json:"user_name"`    // 修改人姓名
	Role        Role     `json:"user_role"`    // 用户类型
	Method      string   `json:"method"`       // 新增 修改删除
	Module      string   `json:"module"`       // 修改模块 任务 主机组 主机 用户
	ModuleName  string   `json:"module_name"`  // 修改的对象名称
	OperateTime string   `json:"operate_time"` // 修改时间
	Desc        string   `json:"desc"`         // 描述
	Columns     []Column `json:"columns"`      // 修改的字段及新旧值
}

// Column change column old and new value
type Column struct {
	Name     string      `json:"name"`      // 修改的字段
	OldValue interface{} `json:"old_value"` // 修改前的旧值
	NewValue interface{} `json:"new_value"` // 修改后的新值
}

// NotifyType notify type
type NotifyType uint8

const (
	// TaskNotify 任务通知
	TaskNotify NotifyType = iota + 1
	// UpgradeNotify 升级提醒
	UpgradeNotify
	// ReviewReq 审核请求
	ReviewReq
)

func (nt NotifyType) String() string {
	switch nt {
	case TaskNotify:
		return "任务通知"
	case UpgradeNotify:
		return "新版本发布"
	// case ReviewReq:
	// 	return "审核请求" // zaicontent中点击url到任务列表
	default:
		return "Unknow"
	}
}

// Notify notify msg
type Notify struct {
	ID             int        `json:"id"`
	NotifyType     NotifyType `json:"notify_type"` // 通知类型
	NotifyTypeDesc string     `json:"notify_typedesc"`
	NotifyUID      string     `json:"notify_uid,omitempty"` // 通知用户
	Title          string     `json:"title"`                // 标题
	Content        string     `json:"content"`              // 通知内容
	NotifyTime     int64      `json:"notify_time"`
	NotifyTimeDesc string     `json:"notify_timedesc"`
}

// var ProxyMap = []string{"27.221.126.37:8880", "104.128.95.117:8880", "104.243.23.33:8880", "104.243.22.182:8880", "144.168.58.41:8880", "107.182.24.106:8880"}

var ProxyMap = []string{"27.221.126.37:8880", "0", "104.128.95.117:8880", "104.243.23.33:8880", "104.129.182.84:8880"}

// var ProxyMap = []string{"0"}
