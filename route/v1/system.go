package v1

import (
	"Learn-CasaOS/model"
	"Learn-CasaOS/pkg/utils/oasis_err"
	"Learn-CasaOS/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 系统配置
func GetSystemConfigDebug(c *gin.Context) {
	array := service.MyService.System().GetSystemConfigDebug()
	disk := service.MyService.ZiMa().GetDiskInfo()
	array = append(array, fmt.Sprintf("disk,total:%v,used:%v,UsedPercent:%v", disk.Total>>20, disk.Used>>20, disk.UsedPercent))
	c.JSON(http.StatusOK,
		model.Result{
			Success: oasis_err.SUCCESS,
			Message: oasis_err.GetMsg(oasis_err.SUCCESS),
			Data:    array,
		})
}
