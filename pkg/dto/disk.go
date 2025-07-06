package dto

import "brabus/pkg/utils"

type Disk struct {
	Space utils.GB         `json:"space"`
	Usage utils.PercentStr `json:"usage"`
}
