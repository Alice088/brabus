package ram

import (
	"github.com/shirou/gopsutil/mem"
	"log"
	"strconv"
)

func GetRAMUsage() string {
	v, err := mem.VirtualMemory()
	if err != nil {
		log.Fatalln("Error during getting RAM usage:", err)
	}

	return strconv.FormatFloat(v.UsedPercent, 'f', 2, 32)
}
