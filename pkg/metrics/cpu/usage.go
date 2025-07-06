package cpu

import (
	"brabus/pkg/dto"
	"brabus/pkg/utils"
	"bufio"
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	prevCPUStats dto.CPUStats
	prevTotal    utils.Percent64
	firstRun     = true
)

func Usage(logger zerolog.Logger) (dto.CPU, error) {
	file, err := os.Open("/proc/stat")
	if err != nil {
		return dto.CPU{}, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logger.Error().Err(err).Msg("Error closing file")
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var cpu dto.CPU
	cpu.Timestamp = time.Now()

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "cpu ") { // Обратите внимание на пробел после cpu!
			fields := strings.Fields(line)
			if len(fields) < 8 {
				continue
			}

			var stats dto.CPUStats
			stats.User = parsePercent(fields[1])
			stats.Nice = parsePercent(fields[2])
			stats.System = parsePercent(fields[3])
			stats.Idle = parsePercent(fields[4])
			stats.IOWait = parsePercent(fields[5])
			stats.Steal = parsePercent(fields[7])

			total := float64(stats.User + stats.Nice + stats.System + stats.Idle + stats.IOWait + stats.Steal)

			if firstRun {
				// Первый запуск - просто сохраняем значения
				prevCPUStats = stats
				prevTotal = total
				firstRun = false
				return dto.CPU{}, fmt.Errorf("first run, collecting baseline")
			} else {
				// Второй и последующие запуски - считаем проценты
				cpu.Average = calculateCPUPercent(stats, prevCPUStats, total, prevTotal)
				prevCPUStats = stats
				prevTotal = total
			}
		} else if strings.HasPrefix(line, "cpu") {
			// Обработка отдельных ядер (если нужно)
			fields := strings.Fields(line)
			if len(fields) < 8 {
				continue
			}

			var stats dto.CPUStats
			stats.User = parsePercent(fields[1])
			stats.Nice = parsePercent(fields[2])
			stats.System = parsePercent(fields[3])
			stats.Idle = parsePercent(fields[4])
			stats.IOWait = parsePercent(fields[5])
			stats.Steal = parsePercent(fields[7])

			cpu.Usage = append(cpu.Usage, stats)
		} else if strings.HasPrefix(line, "model name") {
			cpu.Model = strings.SplitN(line, ":", 2)[1]
			cpu.Model = strings.TrimSpace(cpu.Model)
		}
	}

	cpu.Cores = len(cpu.Usage)
	load1, load5, load15, err := ReadLoadAvg()
	if err == nil {
		cpu.Load1 = load1
		cpu.Load5 = load5
		cpu.Load15 = load15
	}

	return cpu, nil
}

func parsePercent(s string) utils.Percent64 {
	float, _ := strconv.ParseFloat(s, 64)
	return float
}

func ReadLoadAvg() (float64, float64, float64, error) {
	data, err := os.ReadFile("/proc/loadavg")
	if err != nil {
		return 0, 0, 0, err
	}
	fields := strings.Fields(string(data))
	if len(fields) < 3 {
		return 0, 0, 0, fmt.Errorf("invalid loadavg format")
	}
	load1, _ := strconv.ParseFloat(fields[0], 64)
	load5, _ := strconv.ParseFloat(fields[1], 64)
	load15, _ := strconv.ParseFloat(fields[2], 64)
	return load1, load5, load15, nil
}

func calculateCPUPercent(current, prev dto.CPUStats, totalCurrent, totalPrev float64) dto.CPUStats {
	if totalCurrent <= totalPrev {
		return dto.CPUStats{} // избегаем деления на 0 и отрицательных значений
	}

	totalDiff := totalCurrent - totalPrev
	return dto.CPUStats{
		User:   (current.User - prev.User) / totalDiff * 100,
		Nice:   (current.Nice - prev.Nice) / totalDiff * 100,
		System: (current.System - prev.System) / totalDiff * 100,
		Idle:   (current.Idle - prev.Idle) / totalDiff * 100,
		IOWait: (current.IOWait - prev.IOWait) / totalDiff * 100,
		Steal:  (current.Steal - prev.Steal) / totalDiff * 100,
	}
}
