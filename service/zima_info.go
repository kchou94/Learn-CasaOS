package service

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/shirou/gopsutil/v3/disk"
)

type ZiMaService interface {
	GetDiskInfo() *disk.UsageStat
}

type zima struct {
}

// 获取硬盘详情
func (c *zima) GetDiskInfo() *disk.UsageStat {
	path := "/"
	if runtime.GOOS == "windows" {
		path = "C:"
	}
	diskInfo, err := disk.Usage(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	diskInfo.UsedPercent, _ = strconv.ParseFloat(fmt.Sprintf(".1f", diskInfo.UsedPercent), 64)
	diskInfo.InodesUsedPercent, _ = strconv.ParseFloat(fmt.Sprintf(".1f", diskInfo.InodesUsedPercent), 64)
	return diskInfo
}

func NewZiMaService() ZiMaService {
	return &zima{}
}
