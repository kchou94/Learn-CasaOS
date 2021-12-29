package service

import (
	command2 "Learn-CasaOS/pkg/utils/command"
	loger2 "Learn-CasaOS/pkg/utils/loger"
)

type DockerService interface{}

type dockerService struct {
	rootDir string
	log     loger2.Olog
}

func NewDockerService(log loger2.Olog) DockerService {
	return &dockerService{
		rootDir: command2.ExecResultStr(`source ./shell/helper.sh ;GetDockerRootDir`),
		log:     log,
	}
}
