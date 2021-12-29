package main

import (
	"flag"

	"Learn-CasaOS/pkg/config"
	loger2 "Learn-CasaOS/pkg/utils/loger"

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
}
