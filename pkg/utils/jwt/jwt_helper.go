package jwt

import (
	loger2 "Learn-CasaOS/pkg/utils/loger"
	"fmt"
)

func GetToken(username, pwd string) string {
	token, err := GenerateToken(username, pwd)
	if err == nil {
		return token
	} else {
		loger2.NewOLoger().Fatal(fmt.Sprintf("Get Token Fail: %v", err))
		return ""
	}
}
