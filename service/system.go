package service

import (
	"Learn-CasaOS/pkg/config"
	command2 "Learn-CasaOS/pkg/utils/command"
	"Learn-CasaOS/pkg/utils/loger"
)

type SystemService interface {
	GetSystemConfigDebug() []string
}

type systemService struct {
	log loger.OLog
}

func (c *systemService) GetSystemConfigDebug() []string {
	return command2.ExecResultStrArray("source " + config.AppInfo.ProjectPath + "/shell/helper.sh ;GetSysInfo")
}

func NewSystemService(log loger.OLog) SystemService {
	return &systemService{
		log: log,
	}
}
