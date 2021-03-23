package models

type State struct {
	OS     string `json:"os"`
	Uptime string `json:"uptime"`

	CPUName      string `json:"cpuName"`
	CPUCoreCount int32  `json:"cpuCoreCount"`

	MemoryTotal   int32   `json:"memoryTotal"`
	MemoryUsed    int32   `json:"memoryUsed"`
	MemoryPercent float64 `json:"memoryPercent"`

	Temperature float64 `json:"temperature"`
}
