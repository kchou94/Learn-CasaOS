package main

import (
	"flag"

	"Learn-CasaOS/pkg/config"
	"Learn-CasaOS/pkg/sqlite"
	loger2 "Learn-CasaOS/pkg/utils/loger"
	"Learn-CasaOS/route"
	"Learn-CasaOS/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var sqliteDB *gorm.DB

var swagHandler gin.HandlerFunc
var configFlag = flag.String("c", "", "config address")

func init() {
	flag.Parse()
	config.InitSetup(*configFlag)
	loger2.LogSetup()
	sqliteDB = sqlite.GetDb(config.AppInfo.ProjectPath)
	// gredis.GetRedisConn(config.RedisInfo),
	service.MyService = service.NewService(sqliteDB, loger2.NewOLoger())
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
