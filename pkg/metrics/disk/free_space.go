package disk

import (
	"github.com/shirou/gopsutil/disk"
	"strconv"
)

func FreeSpace() string {
	diskUsage, _ := disk.Usage("/")
	return strconv.FormatFloat(float64(diskUsage.Free)/1e9, 'f', 2, 64)
}
