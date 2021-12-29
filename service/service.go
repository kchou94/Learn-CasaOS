package service

var MyService Repository

type Repository interface {
	App() AppService
	DDNS() DDNSService
	User() UserService
	Docker() DockerService
	// Redis() RedisService
	ZeroTier() ZeroTierService
	ZiMa() ZiMaService
	OAPI() OasisService
	Disk() DiskService
	Notify() NotifyServer
	ShareDirectory() ShareDirService
	Task() TaskService
	Rely() RelyService
	System() SystemService
	Shortcuts() ShortcutsService
}
