package dto

type GB = string
type Percent = string

type Disk struct {
	Space GB      `json:"space"`
	Usage Percent `json:"usage"`
}
