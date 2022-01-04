package service

import (
	"Learn-CasaOS/model"
	"Learn-CasaOS/pkg/config"
	httper2 "Learn-CasaOS/pkg/utils/httper"
	json2 "encoding/json"

	"github.com/tidwall/gjson"
)

type CasaService interface {
	GetServerAppInfo(id string) model.ServerAppList
}

type casaService struct {
}

func (o *casaService) GetServerAppInfo(id string) model.ServerAppList {
	head := make(map[string]string)
	head["Authorization"] = GetToken()
	infoS := httper2.Get(config.ServerInfo.ServerApi+"/v2/app/info/"+id, head)
	info := model.ServerAppList{}
	json2.Unmarshal([]byte(gjson.Get(infoS, "data").String()), &info)

	return info
}

func GetToken() string {
	t := make(chan string)
	keyName := "casa_token"

	var auth string
	if result, ok := Cache.Get(keyName); ok {
		auth, ok = result.(string)
		if ok {
			return auth
		}
	}

	go func() {
		str := httper2.Get(config.ServerInfo.ServerApi+"/token", nil)

		t <- gjson.Get(str, "data").String()
	}()
	auth = <-t

	Cache.SetDefault(keyName, auth)
	return auth
}

func NewOasisService() CasaService {
	return &casaService{}
}
