package v1

import "github.com/gin-gonic/gin"

// @Summary 登录zerotier获取token
// @Produce  application/json
// @Accept multipart/form-data
// @Tags zerotier
// @Param username formData string true "User name"
// @Param pwd  formData string true "password"
// @Security ApiKeyAuth
// @Success 200 {string} string "ok"
// @Router /zerotier/login [post]
func ZeroTierGetToken(c *gin.Context) {}
