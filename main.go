package main

import (
	"flag"
	"time"

	"Learn-CasaOS/pkg/config"
	"Learn-CasaOS/pkg/sqlite"
	loger2 "Learn-CasaOS/pkg/utils/loger"
	"Learn-CasaOS/route"
	"Learn-CasaOS/service"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
)

var sqliteDB *gorm.DB

var configFlag = flag.String("c", "", "config address")

var showUserInfo = flag.Bool("show-user-info", false, "show user info")

func init() {
	flag.Parse()
	config.InitSetup(*configFlag)
	config.UpdateSetup()
	loger2.LogSetup()
	sqliteDB = sqlite.GetDb(config.AppInfo.ProjectPath)
	// gredis.GetRedisConn(config.RedisInfo),
	service.MyService = service.NewService(sqliteDB, loger2.NewOLoger())
	service.Cache = cache.New(5*time.Minute, 60*time.Second)
}

// @title Oasis API
// @version 1.0.0
// @contact.name lauren.pan
// @contact.url https://www.zimaboard.com
// @contact.email lauren.pan@icewhale.org
// @description Oasis v1版本api
// @host 192.168.2.114:8089
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /v1

func main() {
	// model.Setup()
	// gredis.Setup()
	r := route.InitRouter(swagHandler)
}
