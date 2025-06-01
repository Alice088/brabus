package ram

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
	"strconv"
)

func Usage() (string, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return "", fmt.Errorf("cannot get RAM usage: %v", err)
	}

	return strconv.FormatFloat(v.UsedPercent, 'f', 2, 32), nil
}
