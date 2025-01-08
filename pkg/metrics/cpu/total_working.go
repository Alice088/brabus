package cpu

import "github.com/c9s/goprocinfo/linux"

func TotalWorking(CPU linux.CPUStat) uint64 {
	return CPU.User + CPU.Nice + CPU.System + CPU.Idle + CPU.IOWait + CPU.IRQ + CPU.SoftIRQ + CPU.Steal
}
