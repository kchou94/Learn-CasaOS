package service

import (
	"Learn-CasaOS/pkg/config"
	command2 "Learn-CasaOS/pkg/utils/command"
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
	GetNet(physics bool) []string
	GetNetState(name string) string
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

// shell脚本参数 {1:虚拟网卡  2:物理网卡}
func (c *zima) GetNet(physics bool) []string {
	t := "1"
	if physics {
		t = "2"
	}
	return command2.ExecResultStrArray("source " + config.AppInfo.ProjectPath + "/shell/helper.sh ;GetNetCard " + t)
}

// shell脚本参数 { 网卡名称 }
func (c *zima) GetNetState(name string) string {
	return command2.ExecResultStr("source " + config.AppInfo.ProjectPath + "/shell/helper.sh ;CatNetCardState " + name)
}

func NewZiMaService() ZiMaService {
	return &zima{}
}
