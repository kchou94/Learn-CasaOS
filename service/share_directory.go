package service

import (
	loger2 "Learn-CasaOS/pkg/utils/loger"

	"gorm.io/gorm"
)

type ShareDirService interface{}

type shareDirService struct {
	db  *gorm.DB
	log loger2.OLog
}

func NewShareDirService(db *gorm.DB, log loger2.OLog) ShareDirService {
	return &shareDirService{
		db:  db,
		log: log,
	}
}
