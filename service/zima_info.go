package service

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
)

type ZiMaService interface {
	GetCpuPercent() float64
	GetCpuCoreNum() int
	GetDiskInfo() *disk.UsageStat
}

type zima struct {
}

// 获取 cpu 占用率
func (z *zima) GetCpuPercent() float64 {
	percent, _ := cpu.Percent(0, false)
	value, _ := strconv.ParseFloat(fmt.Sprintf(".1f", percent[0]), 64)
	return value
}

// 获取 cpu 物理核心数
func (z *zima) GetCpuCoreNum() int {
	count, _ := cpu.Counts(false)
	return count
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
