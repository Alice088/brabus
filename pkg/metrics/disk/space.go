package disk

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"strconv"
)

func Space() (string, error) {
	diskUsage, err := disk.Usage("/")

	if err != nil {
		return "", fmt.Errorf("cannot get disk usage: %v", err)
	}

	return strconv.FormatFloat(float64(diskUsage.Free)/1e9, 'f', 2, 64), nil
}
