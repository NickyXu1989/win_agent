package utils

import (
	"container/list"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"sync"
	"time"
)

type cpuMetric struct {
	cpuTime cpu.TimesStat
	execTime time.Time
}

type CpuMetricsHistory struct {
	lock sync.Mutex
	cpuTimesHistory *list.List
}

var instance *CpuMetricsHistory

func NewCpuMetricsHistory() *CpuMetricsHistory {
	if instance == nil {
		instance = &CpuMetricsHistory{
			lock:            sync.Mutex{},
			cpuTimesHistory: list.New(),
		}
	}
	return instance
}

func GetCpuMetricsHistoryInstance() *CpuMetricsHistory {
	return instance
}

func (c *CpuMetricsHistory) GatherCpuMetric() {
	//fmt.Println("gather cpu metric")
	c.lock.Lock()
	defer c.lock.Unlock()

	cpuTime := GetCpuTimes()
	cpuMetric := cpuMetric{
		cpuTime:  cpuTime,
		execTime: time.Now(),
	}

	c.cpuTimesHistory.PushBack(cpuMetric)

	if c.cpuTimesHistory.Len() > 100 {
		de := c.cpuTimesHistory.Front()
		c.cpuTimesHistory.Remove(de)
	}
}

func (c *CpuMetricsHistory) PrintAllCpuMetrics() {
	for item := c.cpuTimesHistory.Front();nil != item ;item = item.Next() {
		fmt.Println(item.Value)
	}
	fmt.Println(c.cpuTimesHistory.Len())
}


//获取cpu百分比
func (c *CpuMetricsHistory) GetLatestCpuPercent() float64 {
	c.lock.Lock()
	defer c.lock.Unlock()

	lastCpuMetric := c.cpuTimesHistory.Back().Value.(cpuMetric)
	lastCpuStat := lastCpuMetric.cpuTime
	currentCpuStat := GetCpuTimes()

	currentCpuSum := currentCpuStat.User + currentCpuStat.System + currentCpuStat.Idle + currentCpuStat.Nice + currentCpuStat.Iowait + currentCpuStat.Softirq + currentCpuStat.Irq + currentCpuStat.Steal + currentCpuStat.GuestNice + currentCpuStat.Guest
	lastCpuCountSum := lastCpuStat.User + lastCpuStat.System + lastCpuStat.Idle + lastCpuStat.Nice + lastCpuStat.Iowait + lastCpuStat.Softirq + lastCpuStat.Irq + lastCpuStat.Steal + lastCpuStat.GuestNice + lastCpuStat.Guest

	deltaCpuSum := currentCpuSum - lastCpuCountSum
	deltaCpuIdle := currentCpuStat.Idle - lastCpuStat.Idle
	var cpuBusyPct float64
	if 0 != deltaCpuSum {
		cpuBusyPct = (1 -(deltaCpuIdle/deltaCpuSum)) * 100
	} else {
		cpuBusyPct = 0
	}

	return cpuBusyPct

}
