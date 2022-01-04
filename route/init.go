package route

import (
	"Learn-CasaOS/model"
	"Learn-CasaOS/model/system_app"
	"Learn-CasaOS/pkg/config"
	"Learn-CasaOS/pkg/docker"
	"Learn-CasaOS/pkg/utils/file"
	"Learn-CasaOS/pkg/utils/port"
	"Learn-CasaOS/service"
	"encoding/json"
	"encoding/xml"
	"strconv"
	"time"
)

func InitFunction() {
	go checkSystemApp()
	Update2_3()
	CheckSerialDiskMount()
}

var syncIsExistence = false

func installSyncthing(appId string) {

	appInfo := service.MyService.OAPI().GetServerAppInfo(appId)
	dockerImage := appInfo.Image
	m := model.CustomizationPostData{}
	var dockerImageVersion string

	if len(appInfo.ImageVersion) == 0 {
		dockerImageVersion = "latest"
	}

	if appInfo.NetworkModel != "host" {
		for i := 0; i < len(appInfo.Ports); i++ {
			if p, _ := strconv.Atoi(appInfo.Ports[i].ContainerPort); port.IsPortAvailable(p, appInfo.Ports[i].Protocol) {
				appInfo.Ports[i].CommendPort = strconv.Itoa(p)
			} else {
				if appInfo.Ports[i].Protocol == "tcp" {
					if p, err := port.GetAvailablePort("tcp"); err == nil {
						appInfo.Ports[i].CommendPort = strconv.Itoa(p)
					}
				} else if appInfo.Ports[i].Protocol == "upd" {
					if p, err := port.GetAvailablePort("udp"); err == nil {
						appInfo.Ports[i].CommendPort = strconv.Itoa(p)
					}
				}
			}

			if appInfo.Ports[i].Type == 0 {
				appInfo.PortMap = appInfo.Ports[i].CommendPort
			}
		}
	}
}

// check if the system application is installed
func checkSystemApp() {
	list := service.MyService.App().GetSystemAppList()
	for _, v := range *list {
		if v.Image == "linuxserver/syncthing" {
			if v.State != "running" {
				service.MyService.Docker().DockerContainerStart(v.CustomId)
			}
			syncIsExistence = true
			if config.SystemConfigInfo.SyncPort != v.Port {
				config.SystemConfigInfo.SyncPort = v.Port
			}
			var paths []model.PathMap
			json.Unmarshal([]byte(v.Volumes), &paths)
			path := ""
			for _, i := range paths {
				if i.ContainerPath == "/config" {
					path = docker.GetDir(v.CustomId, i.Path) + "config.xml"
					for i := 0; i < 10; i++ {
						if file.CheckNotExist(path) {
							time.Sleep(1 * time.Second)
						} else {
							break
						}
					}
					break
				}
			}
			content := file.ReadFullFile(path)
			syncConfig := &system_app.SyncConfig{}
			xml.Unmarshal(content, &syncConfig)
			config.SystemConfigInfo.SyncKey = syncConfig.Key
		}
	}
	if !syncIsExistence {
		installSyncthing("74")
	}
}
