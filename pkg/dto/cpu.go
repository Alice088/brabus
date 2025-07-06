package dto

import (
	"brabus/pkg/utils"
	"time"
)

type CPUStats struct {
	User   utils.Percent64 `json:"user"`
	System utils.Percent64 `json:"system"`
	Idle   utils.Percent64 `json:"idle"`
	IOWait utils.Percent64 `json:"iowait"`
	Steal  utils.Percent64 `json:"steal"`
	Nice   utils.Percent64 `json:"nice"`
}

type CPU struct {
	Usage     []CPUStats `json:"usage"`
	Average   CPUStats   `json:"average"`
	Cores     int        `json:"cores"`
	Model     string     `json:"model"`
	Load1     float64    `json:"load1"`
	Load5     float64    `json:"load5"`
	Load15    float64    `json:"load15"`
	Timestamp time.Time  `json:"timestamp"`
}
