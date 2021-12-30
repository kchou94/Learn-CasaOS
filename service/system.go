package service

import (
	"Learn-CasaOS/pkg/config"
	command2 "Learn-CasaOS/pkg/utils/command"
)

type SystemService interface {
	GetSystemConfigDebug() []string
}

type systemService struct{}

func (c *systemService) GetSystemConfigDebug() []string {
	return command2.ExecResultStrArray("source " + config.AppInfo.ProjectPath + "/shell/helper.sh ;GetSysInfo")
}

func NewSystemService() SystemService {
	return &systemService{}
}
