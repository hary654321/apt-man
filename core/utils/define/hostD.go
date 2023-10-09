package define

// HostGroup define hostgroup
type HostGroup struct {
	HostsID     []string `json:"addrs" comment:"WorkerIDs"` // 主机host
	CreateByUID string   `json:"create_byuid"`              // 创建人ID
	CreateBy    string   `json:"create_by"`                 // 创建人ID
	Common
}

// CreateHostGroup new hostgroup
type CreateHostGroup struct {
	Name    string   `json:"name" binding:"required,max=30"`
	HostsID []string `json:"addrs"` // 主机host
	Remark  string   `json:"remark" binding:"max=100"`
}

// ChangeHostGroup new hostgroup
type ChangeHostGroup struct {
	ID      string   `json:"id" binding:"required"`
	HostsID []string `json:"addrs"` // 主机host
	Remark  string   `json:"remark" binding:"max=100"`
}

// Host worker host
type Host struct {
	ID                 string   `json:"id" comment:"ID"`
	Status             int      `json:"status" comment:"状态"`
	Addr               string   `json:"addr" `
	Ip                 string   `json:"ip" comment:"客户端ip"`
	HostName           string   `json:"hostname" comment:"主机名称"`
	Online             bool     `json:"online"`
	Weight             int      `json:"weight"`
	SshPort            int      `json:"sshPort" comment:"ssh端口"`
	ServicePort        int      `json:"servicePort"`
	RunningTasks       []string `json:"running_tasks"`
	Version            string   `json:"version"`
	SshUser            string   `json:"sshUser" comment:"ssh用户"`
	SshPwd             string   `json:"sshPwd"`
	Stop               bool     `json:"stop" comment:"暂停调度"`
	LastUpdateTimeUnix int64    `json:"last_updatetimeunix"`
	LastUpdateTime     string   `json:"last_updatetime" comment:"更新时间"`
	Remark             string   `json:"remark"`
}

type HostGorm struct {
	ID          string `gorm:"column:id" json:"id" comment:"ID"`
	Ip          string `gorm:"column:ip" json:"ip" comment:"客户端ip"`
	HostName    string `gorm:"column:hostname" json:"hostname"`
	Weight      int    `gorm:"column:weight" json:"weight"`
	SshPort     int    `gorm:"column:sshPort" json:"sshPort"`
	ServicePort int    `gorm:"column:servicePort" json:"servicePort"`
	Version     string `gorm:"column:version" json:"version"`
	SshUser     string `gorm:"column:sshUser" json:"sshUser"`
	SshPwd      string `gorm:"column:sshPwd" json:"sshPwd"`
	Remark      string `gorm:"column:remark" json:"remark"`
}
