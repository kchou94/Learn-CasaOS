package service

import "gorm.io/gorm"

type NotifyServer interface{}

type notifyServer struct {
	db *gorm.DB
}

func NewNotifyService(db *gorm.DB) NotifyServer {
	return &notifyServer{
		db: db,
	}
}
