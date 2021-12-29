package service

import (
	loger2 "Learn-CasaOS/pkg/utils/loger"

	"gorm.io/gorm"
)

type TaskService interface{}

type taskService struct {
	db  *gorm.DB
	log loger2.OLog
}

func NewTaskService(db *gorm.DB, log loger2.OLog) TaskService {
	return &taskService{
		db:  db,
		log: log,
	}
}
