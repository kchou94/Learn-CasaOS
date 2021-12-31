package service

import (
	command2 "Learn-CasaOS/pkg/utils/command"
	loger2 "Learn-CasaOS/pkg/utils/loger"
	"context"

	"github.com/docker/docker/api/types"
	client2 "github.com/docker/docker/client"
)

type DockerService interface {
	DockerContainerStart(name string) error
}

type dockerService struct {
	rootDir string
	log     loger2.OLog
}

//启动容器
func (d *dockerService) DockerContainerStart(name string) error {
	cli, err := client2.NewClientWithOpts(client2.FromEnv)
	if err != nil {
		return err
	}
	defer cli.Close()
	err = cli.ContainerStart(context.Background(), name, types.ContainerStartOptions{})
	if err != nil {
		return err
	}
	return nil
}

func NewDockerService(log loger2.OLog) DockerService {
	return &dockerService{
		rootDir: command2.ExecResultStr(`source ./shell/helper.sh ;GetDockerRootDir`),
		log:     log,
	}
}
