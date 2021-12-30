package sqlite

import (
	"fmt"
	"time"

	"Learn-CasaOS/pkg/utils/file"
	model2 "Learn-CasaOS/service/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var gdb *gorm.DB

func GetDb(projectPath string) *gorm.DB {
	if gdb != nil {
		return gdb
	}

	err := file.IsNotExistMkDir(projectPath + "/db/")
	if err != nil {
		fmt.Println(err)
	}
	db, err := gorm.Open(sqlite.Open(projectPath+"/db/CasaOS.db"), &gorm.Config{})
	c, _ := db.DB()
	c.SetMaxIdleConns(10)
	c.SetMaxOpenConns(100)
	c.SetConnMaxIdleTime(time.Second * 1000)
	if err != nil {
		fmt.Println("连接数据失败!")
		panic("数据连接失败")
	}

	gdb = db
	err = gdb.AutoMigrate(&model2.TaskDBModel{}, &model2.AppNotify{}, &model2.AppListDBModel{})
	if err != nil {
		fmt.Println("检查和创建数据库出错", err)
	}
	return gdb
}
