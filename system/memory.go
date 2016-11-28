package system

import (
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/host"
)

type MemoryInfo struct {
	Total uint64 `json:"total"`
	Free  uint64 `json:"free"`
	UsedPercent float64 `json:"used_percent"`
}

type SystemInfo struct {
	Hostname             string `json:"hostname"`
	Uptime               uint64 `json:"uptime"`
	BootTime             uint64 `json:"bootTime"`
}

func GetMemoryInfo()  *MemoryInfo{
	v, _ := mem.VirtualMemory()
	return &MemoryInfo{
		Total:v.Total,
		Free:v.Free,
		UsedPercent: v.UsedPercent,
	}
}

func GetSystemInfo() *SystemInfo{
	h, _ := host.Info()
	return &SystemInfo{
		Hostname: h.Hostname,
		BootTime: h.BootTime,
		Uptime: h.Uptime,
	}
}