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
}

// 服务配置
type ServerModel struct {
	HttpPort  string
	RunMode   string
	ServerApi string
}

// 服务配置
type AppModel struct {
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
type ZerotierModel struct {
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
	SerachSwitch    bool   `json:"search_switch"`    // 搜索开关
	SearchEngine    string `json:"search_engine"`    // 搜索引擎
	ShortcutsSwitch bool   `json:"shortcuts_switch"` // 快捷键开关
	WidgetsSwitch   bool   `json:"widgets_switch"`   // 插件开关
	BackgroundType  string `json:"background_type"`  // 背景类型
	Background      string `json:"background"`       // 背景图片
	AutoUpdate      bool   `json:"auto_update"`      // 自动更新
}
