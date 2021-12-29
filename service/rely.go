package service

import (
	loger2 "Learn-CasaOS/pkg/utils/loger"

	"gorm.io/gorm"
)

type RelyService interface{}

type relyService struct {
	db  *gorm.DB
	log loger2.OLog
}

func NewRelyService(db *gorm.DB, log loger2.OLog) RelyService {
	return &relyService{db: db, log: log}
}
