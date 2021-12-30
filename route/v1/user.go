package v1

import (
	"Learn-CasaOS/model"
	"Learn-CasaOS/pkg/config"
	jwt2 "Learn-CasaOS/pkg/utils/jwt"
	oasis_err2 "Learn-CasaOS/pkg/utils/oasis_err"
	"Learn-CasaOS/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var user_service service.UserService

func init() {
	user_service = service.NewUserService()
}

// @Summary 设置用户名和密码
// @Produce  application/json
// @Accept multipart/form-data
// @Tags user
// @Param username formData string true "User name"
// @Param pwd  formData string true "password"
// @Security ApiKeyAuth
// @Success 200 {string} string "ok"
// @Router /user/setusernamepwd [post]
func Set_Name_Pwd(c *gin.Context) {
	// json := make(map[string]string)
	// c.BindJSON(&json)
	username := c.PostForm("username")
	pwd := c.PostForm("pwd")
	// 老用户名是否存在即新用户名和密码的验证
	if len(config.UserInfo.UserName) > 0 || len(username) == 0 || len(pwd) == 0 {
		c.JSON(http.StatusOK,
			model.Result{
				Success: oasis_err2.ERROR,
				Message: oasis_err2.GetMsg(oasis_err2.INVALID_PARAMS),
			})
		return
	}
	// 开始设置
	err := user_service.SetUser(username, pwd, "", "", "")
	if err != nil {
		c.JSON(http.StatusOK,
			model.Result{
				Success: oasis_err2.SUCCESS,
				Message: fmt.Sprintf("%v", err),
			})
		return
	} else {
		c.JSON(http.StatusOK,
			model.Result{
				Success: oasis_err2.SUCCESS,
				Message: oasis_err2.GetMsg(oasis_err2.SUCCESS),
			})
		return
	}
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
		// user_service.SetUser("", "", token, "", "")
		c.JSON(http.StatusOK,
			model.Result{
				Success: oasis_err2.SUCCESS,
				Message: oasis_err2.GetMsg(oasis_err2.SUCCESS),
				Data:    token,
			})
		return
	}
	c.JSON(http.StatusOK,
		model.Result{
			Success: oasis_err2.ERROR,
			Message: oasis_err2.GetMsg(oasis_err2.ERROR),
		})
}
