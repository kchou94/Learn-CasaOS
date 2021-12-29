package service

import (
	"net/http"
	"time"
)

type ZeroTierService interface{}

type zerotierStruct struct {
}

func NewZeroTierService() ZeroTierService {
	// 初始化 client
	client := http.Client{Timeout: 30 * time.Second, CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse // 防止重定向
	},
	}
	return &zerotierStruct{}
}
