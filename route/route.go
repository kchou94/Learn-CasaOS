package route

import (
	"Learn-CasaOS/middleware"
	"Learn-CasaOS/pkg/config"
	v1 "Learn-CasaOS/route/v1"
	"Learn-CasaOS/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

var swagHandler gin.HandlerFunc

func InitRouter(swagHandler gin.HandlerFunc) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	gin.SetMode(config.ServerInfo.RunMode)

	r.StaticFS("/ui", http.FS(web.Static))
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "ui/")
	})

	if swagHandler != nil {
		r.GET("/swagger/*any", swagHandler)
	}

	// 登录
	r.POST("/v1/user/login", v1.Login)

	r.GET("/debug", v1.GetSystemConfigDebug)

	return r
}
