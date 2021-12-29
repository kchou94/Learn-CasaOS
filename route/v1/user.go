package v1

import (
	"Learn-CasaOS/model"
	jwt2 "Learn-CasaOS/pkg/utils/jwt"
	oasis_err2 "Learn-CasaOS/pkg/utils/oasis_err"
	"Learn-CasaOS/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var user_service service.UserService

func init() {
	user_service = service.NewUserService()
}

// @Summary 登录
// @Produce  application/json
// @Accept multipart/form-data
// @Tags user
// @Param username formData string true "User name"
// @Param pwd  formData string true "password"
// @Success 200 {string} string "ok"
// @Router /user/login [post]
func Login(c *gin.Context) {
	username := c.PostForm("username")
	pwd := c.PostForm("pwd")

	// 检查参数是否正确
	if len(username) == 0 || len(pwd) == 0 {
		c.JSON(http.StatusOK,
			model.Result{
				Success: oasis_err2.ERROR,
				Message: oasis_err2.GetMsg(oasis_err2.INVALID_PARAMS),
			})
		return
	}

	// if config.UserInfo.UserName == username && config.UserInfo.PWD == pwd {
	if username == "admin" && pwd == "admin" {
		token := jwt2.GetToken(username, pwd)

	}
}
