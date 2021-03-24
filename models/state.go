package models

type State struct {
	OS     string `json:"os"`
	Uptime string `json:"uptime"`

	CPUName      string `json:"cpuName"`
	CPUCoreCount int32  `json:"cpuCoreCount"`

	MemoryTotal   uint32  `json:"memoryTotal"`
	MemoryUsed    uint32  `json:"memoryUsed"`
	MemoryPercent float64 `json:"memoryPercent"`

	DiskTotal       uint32  `json:"total"`
	DiskFree        uint32  `json:"free"`
	DiskUsed        uint32  `json:"used"`
	DiskUsedPercent float64 `json:"usedPercent"`

	Temperature float64 `json:"temperature"`
}
