package service

import loger2 "Learn-CasaOS/pkg/utils/loger"

type DiskService interface{}

type diskService struct {
	log loger2.OLog
}

func NewDiskService(log loger2.OLog) DiskService {
	return &diskService{
		log: log,
	}
}
