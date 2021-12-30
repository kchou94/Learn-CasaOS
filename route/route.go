package route

import (
	"Learn-CasaOS/middleware"
	"Learn-CasaOS/pkg/config"
	jwt2 "Learn-CasaOS/pkg/utils/jwt"
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

	v1Group := r.Group("/v1")

	v1Group.Use(jwt2.JWT(swagHandler))
	{
		v1UserGroup := v1Group.Group("/user")
		v1UserGroup.Use()
		{
			// 设置用户
			v1UserGroup.POST("/setusernamepwd", v1.Set_Name_Pwd)
			// 修改头像
			v1UserGroup.POST("/changhead", v1.Up_Load_Head)
			// 修改用户名
			v1UserGroup.POST("/changusername", v1.Chang_User_Name)
		}
	}

	return r
}
