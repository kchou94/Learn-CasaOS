package service

import (
	loger2 "Learn-CasaOS/pkg/utils/loger"
	"Learn-CasaOS/service/model"

	"gorm.io/gorm"
)

type ddnsStruct struct {
	db  *gorm.DB
	log loger2.OLog
}

type DDNSService interface {
	IsExist(t int, domain string, host string) bool
	GetExternalIP() (string, string)
	GetConfigList() *[]model.DDNSList
	DeleteConfig(id uint) bool
	GetType(name string) (uint, string)
	SaveConfig(model model.DDNSUpdateDBModel) error
}

// 判断当前添加的是否存在
func (d *ddnsStruct) IsExist(t int, domain string, host string) bool {
	return true
}

// 前台获取已配置的ddns列表
func (d *ddnsStruct) GetConfigList() *[]model.DDNSList {
	var s []model.DDNSList
	return &s
}

func (d *ddnsStruct) DeleteConfig(id uint) bool {
	return true
}

func (d *ddnsStruct) GetExternalIP() (string, string) {
	return "", ""
}

// 根据名称获取类型
func (d *ddnsStruct) GetType(name string) (uint, string) {
	return 0, ""
}

//保存配置到数据库
func (d *ddnsStruct) SaveConfig(model model.DDNSUpdateDBModel) error {
	return nil
}

func NewDDNSService(db *gorm.DB, log loger2.OLog) DDNSService {
	return &ddnsStruct{db, log}
}
