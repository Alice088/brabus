package DTO

type Metrics struct {
	CPUUsage  []string `json:"cpu_usage"`
	RAMUsage  string   `json:"ram_usage"`
	DiskSpace string   `json:"disk_space"`
	DiskUsage string   `json:"disk_usage"`
}
