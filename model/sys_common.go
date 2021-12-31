package model

import "time"

// 系统配置
type SysInfoModel struct {
	Name string // 系统名称
}

// 用户相关
type UserModel struct {
	UserName    string
	PWD         string
	Token       string
	Head        string
	Email       string
	Description string
	Initialized bool //是否初始化
}

// 服务配置
type ServerModel struct {
	HttpPort    string
	RunMode     string
	ServerApi   string
	LockAccount bool //是否锁定账号
}

// 服务配置
type APPModel struct {
	LogSavePath    string
	LogSaveName    string
	LogFileExt     string
	DateStrFormat  string
	DateTimeFormat string
	TimeFormat     string
	DateFormat     string
	ProjectPath    string
}

// 公共返回模型
type Result struct {
	Success int         `json:"success" example:"200"`
	Message string      `json:"message" example:"ok"`
	Data    interface{} `json:"data" example:"返回结果"`
}

// zerotier
type ZeroTierModel struct {
	UserName string
	PWD      string
	Token    string
}

// redis config
type RedisModel struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type SystemConfig struct {
	ConfigStr  string `json:"config_str"`  // 配置字符串
	WidgetList string `json:"widget_list"` // 小部件列表
	ConfigPath string `json:"config_path"` // 配置文件路径
	SyncPort   string `json:"sync_port"`   // 同步端口
	SyncKey    string `json:"sync_key"`    // 同步密钥
}

type CasaOSGlobalVariables struct {
	AppChange bool // app是否更新
}
