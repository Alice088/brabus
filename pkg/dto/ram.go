package dto

import "brabus/pkg/utils"

type RAM struct {
	Usage utils.PercentStr `json:"usage"`
}
