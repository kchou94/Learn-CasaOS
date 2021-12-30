package jwt

import (
	"Learn-CasaOS/model"
	loger2 "Learn-CasaOS/pkg/utils/loger"
	oasis_err2 "Learn-CasaOS/pkg/utils/oasis_err"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT(swagHandler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := oasis_err2.SUCCESS
		token := c.GetHeader("Authorization")
		if len(token) == 0 {
			token = c.Query("token")
		}
		if token == "" {
			code = oasis_err2.INVALID_PARAMS
		}
		if swagHandler == nil {
			claims, err := ParseToken(token)
			if err != nil {
				code = oasis_err2.ERROR_AUTH_TOKEN
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = oasis_err2.ERROR_AUTH_TOKEN
			}
		}

		if code != oasis_err2.SUCCESS {
			c.JSON(http.StatusOK,
				model.Result{
					Success: code,
					Message: oasis_err2.GetMsg(code),
				})
			c.Abort()
			return
		}
		c.Next()
	}
}

func GetToken(username, pwd string) string {
	token, err := GenerateToken(username, pwd)
	if err == nil {
		return token
	} else {
		loger2.NewOLoger().Fatal(fmt.Sprintf("Get Token Fail: %v", err))
		return ""
	}
}
