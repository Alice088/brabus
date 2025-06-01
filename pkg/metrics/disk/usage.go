package disk

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"strconv"
)

func Usage() (string, error) {
	diskUsage, err := disk.Usage("/")

	if err != nil {
		return "", fmt.Errorf("cannot get disk usage: %v", err)
	}

	return strconv.FormatFloat(diskUsage.UsedPercent, 'f', 2, 32), nil
}
