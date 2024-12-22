package ram

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
	"strconv"
)

func GetRAMUsage() (string, error) {
	v, err := mem.VirtualMemory()
	if err == nil {
		return strconv.FormatFloat(v.UsedPercent, 'f', 2, 32), nil
	}

	fmt.Println("Error during getting RAM usage:", err)
	return ``, err
}
