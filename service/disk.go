package service

import (
	loger2 "Learn-CasaOS/pkg/utils/loger"

	"gorm.io/gorm"
)

type DiskService interface{}

type diskService struct {
	log loger2.OLog
	db  *gorm.DB
}

func NewDiskService(log loger2.OLog, db *gorm.DB) DiskService {
	return &diskService{
		log: log,
		db:  db,
	}
}
