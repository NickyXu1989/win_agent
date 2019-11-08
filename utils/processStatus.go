package utils

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
)

type WinProcess struct {
	Pid int32
	Name string
	String string
	CPUPercent float64
	CreateTime int64
	Cmdline string
	MemoryPercent float32
	MemoryStat *process.MemoryInfoStat
}

func GetProcessStatus() []WinProcess {
	processes, err := process.Processes()
	if err != nil {
		fmt.Println(err.Error())
	}

	var winProcesses []WinProcess
	for _, process := range processes {
		pid := process.Pid
		name,_ := process.Name()
		str := process.String()
		cpuPct,_ := process.CPUPercent()
		createTime,_ := process.CreateTime()
		ioCouters,_ :=process.IOCounters()
		fmt.Println(ioCouters)
		cmdLine,_ := process.Cmdline()
		memoryPercent, _ := process.MemoryPercent()
		memoryStat,_ := process.MemoryInfo()
		winProcess := WinProcess{
			Pid: pid,
			Name: name,
			String: str,
			CPUPercent:cpuPct,
			CreateTime:createTime,
			Cmdline:cmdLine,
			MemoryPercent:memoryPercent,
			MemoryStat: memoryStat,
		}
		winProcesses = append(winProcesses, winProcess)
	}
	return winProcesses
}
