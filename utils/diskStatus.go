package utils

import (
	"github.com/shirou/gopsutil/disk"
)


type DiskIOMetric struct {
	DiskName string
	ReadCount float64
	MergedReadCount float64
	WriteCount float64
	MergedWriteCount float64
	ReadBytes float64
	WriteBytes float64
	ReadTime float64
	WriteTime float64
	IopsInProgress float64
	IoTime float64
	WeightedIO float64
}

func GetDiskUsagePct()  []map[string]interface{}{
	var diskMetrics []map[string]interface{}
	drivers, _ := disk.Partitions(true)
	for _,driver:= range drivers {
		du,_ := disk.Usage(driver.Device)
		diskMetrics = append(diskMetrics, map[string]interface{}{
			"path": driver.Device,
			"value": du.UsedPercent,
		})
	}
	return diskMetrics
}

func GetDiskIOMetrics() []DiskIOMetric {
	//get all disks
	drivers, _ := disk.Partitions(true)
	var diskIOMetrics []DiskIOMetric
	for _,driver:= range drivers {
		//get disk io status
		ds, _ := disk.IOCounters(driver.Device)
		for k,v := range ds {
			diskIOMetrics = append(diskIOMetrics, DiskIOMetric{
				DiskName:         k,
				ReadCount:        float64(v.ReadCount),
				MergedReadCount:  float64(v.MergedReadCount),
				WriteCount:       float64(v.WriteCount),
				MergedWriteCount: float64(v.MergedWriteCount),
				ReadBytes:        float64(v.ReadBytes),
				WriteBytes:       float64(v.WriteBytes),
				ReadTime:         float64(v.ReadTime),
				WriteTime:        float64(v.WriteTime),
				IopsInProgress:   float64(v.IopsInProgress),
				IoTime:           float64(v.IoTime),
				WeightedIO:       float64(v.WeightedIO),
			})
		}
	}
	return diskIOMetrics
}