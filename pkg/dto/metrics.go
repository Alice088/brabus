package dto

type Metrics struct {
	CPU  CPU  `json:"cpu"`
	RAM  RAM  `json:"ram"`
	Disk Disk `json:"disk"`
}
