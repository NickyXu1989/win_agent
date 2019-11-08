package utils

import (
	"github.com/shirou/gopsutil/cpu"
)

func GetCpuTimes() cpu.TimesStat {
	cpuTimes, _ := cpu.Times(false)
	return cpuTimes[0]
}