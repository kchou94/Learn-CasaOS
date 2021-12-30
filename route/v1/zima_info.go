package v1

import (
	"Learn-CasaOS/model"
	oasis_err2 "Learn-CasaOS/pkg/utils/oasis_err"
	"Learn-CasaOS/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 获取cpu信息
// @Produce  application/json
// @Accept application/json
// @Tags zima
// @Security ApiKeyAuth
// @Success 200 {string} string "ok"
// @Router /zima/getcpuinfo [get]
func CpuInfo(c *gin.Context) {
	// 检查参数是否正确
	cpu := service.MyService.ZiMa().GetCpuPercent()
	num := service.MyService.ZiMa().GetCpuCoreNum()
	data := make(map[string]interface{})
	data["percent"] = cpu
	data["num"] = num
	c.JSON(http.StatusOK,
		model.Result{
			Success: oasis_err2.SUCCESS,
			Message: oasis_err2.GetMsg(oasis_err2.SUCCESS),
			Data:    data,
		})
}
