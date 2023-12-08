package config

import (
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

var (
	// CoreConf crocodile conf
	CoreConf *coreConf
)

// Init Config
func Init(conf string) {
	_, err := toml.DecodeFile(conf, &CoreConf)
	if err != nil {
		fmt.Printf("Err %v", err)
		os.Exit(1)
	}
}

type coreConf struct {
	SecretToken string
	BasicAuth   string
	Sso         Sso
	Log         Log
	Cert        Cert
	Server      Server
	Client      Client
	Notify      Notify
}
type Sso struct {
	SsoPort  int
	SsoIp    string
	Secm     string
	Username string
	Password string
}

// Log Config
type Log struct {
	LogPath    string
	MaxSize    int
	Compress   bool
	MaxAge     int
	MaxBackups int
	LogLevel   string
	Format     string
}

// Cert tls cert
type Cert struct {
	Enable   bool
	CertFile string
	KeyFile  string
}

// Server crocodile server config
type Server struct {
	LogPath     string
	Port        int
	Sleep       time.Duration
	Run         bool
	MaxHTTPTime duration
	DB          db
	Redis       redis
}

type db struct {
	Drivename    string
	TablePrefix  string
	Dsn          string
	MaxIdle      int
	MaxConn      int
	MaxQueryTime duration
}

type redis struct {
	Addr     string
	PassWord string
}

// Client crocodile client config
type Client struct {
	Port        int
	ServerAddrs []string
	ServerPort  int
	HostGroup   string
	Weight      int
	Remark      string
}

type duration struct {
	time.Duration
}

// UnmarshalText parse 10s to time.Time
func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

// Notify send msg to user
type Notify struct {
	Email    email
	DingDing dingding
	Slack    slack
	Telegram telegram
	WeChat   wechat
	WebHook  webhook
}

type email struct {
	Enable     bool
	SMTPHost   string
	Port       int
	UserName   string
	Password   string
	From       string
	TLS        bool
	Anonymous  bool
	SkipVerify bool
}

type dingding struct {
	Enable      bool
	WebHook     string
	SecureLevel int
	Secret      string
}

type slack struct {
	Enable  bool
	WebHook string
}

type telegram struct {
	Enable   bool
	BotToken string
}

type wechat struct {
	Enable      bool
	CropID      string
	AgentID     int
	AgentSecret string
}

type webhook struct {
	Enable     bool
	WebHookURL string
}
