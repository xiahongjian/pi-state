package service

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/xiahongjian/pi-status/models"
)

func GetState() *models.State {

	h, _ := host.Info()
	m, _ := mem.VirtualMemory()
	c, _ := cpu.Info()

	info := &models.State{
		OS:     fmt.Sprintf("%s %s", h.Platform, h.PlatformVersion),
		Uptime: formatUptime(h.Uptime),

		CPUName:      c[0].ModelName,
		CPUCoreCount: c[0].Cores,

		MemoryTotal:   byteToMB(m.Total),
		MemoryUsed:    byteToMB(m.Used),
		MemoryPercent: float32(m.UsedPercent),
	}
	return info
}

func byteToMB(byteNumber uint64) int32 {
	result := int32(byteNumber >> 20)
	return result
}

func formatUptime(uptime uint64) string {

	str := ""

	totalMinis := uptime / 60
	totalHours := totalMinis / 60
	totalDays := totalHours / 24

	if totalDays > 0 {
		str += fmt.Sprintf("%d天", totalDays)
	}
	hours := totalHours - (totalDays * 24)
	if hours > 0 || totalDays > 0 {
		str += fmt.Sprintf("%d小时", hours)
	}
	minutes := totalMinis - (totalDays*24+hours)*60
	str += fmt.Sprintf("%d分钟", minutes)
	return str
}
