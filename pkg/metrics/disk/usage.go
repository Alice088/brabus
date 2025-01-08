package disk

import (
	"github.com/shirou/gopsutil/disk"
	"strconv"
)

func Usage() string {
	diskUsage, _ := disk.Usage("/")
	return strconv.FormatFloat(diskUsage.UsedPercent, 'f', 2, 32)
}
