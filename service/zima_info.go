package service

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

type ZiMaService interface {
	GetCpuPercent() float64
	GetCpuCoreNum() int
	GetMemInfo() *mem.VirtualMemoryStat
	GetDiskInfo() *disk.UsageStat
	GetNetInfo() []net.IOCountersStat
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

// 获取内存详情
func (z *zima) GetMemInfo() *mem.VirtualMemoryStat {
	memInfo, _ := mem.VirtualMemory()
	return memInfo
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

// 网络信息
func (c *zima) GetNetInfo() []net.IOCountersStat {
	netInfo, _ := net.IOCounters(true)
	return netInfo
}

func NewZiMaService() ZiMaService {
	return &zima{}
}
