package utils

import "github.com/shirou/gopsutil/mem"

func GetMemPct() float64{
	v, _ := mem.VirtualMemory()
	memUsed := v.UsedPercent
	return memUsed
}
