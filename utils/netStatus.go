package utils

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/shirou/gopsutil/net"
)

type NetMetrics struct {
	Name string
	BytesSent float64
	BytesRecv float64
	PacketsSent float64
	PacketsRecv float64
	Errin float64
	Errout float64
	Dropin float64
	Dropout float64
	Fifoin float64
	Fifoout float64
}

func GetNetStatus() []NetMetrics {
	tmpSet := mapset.NewSet()
	var netMetrics []NetMetrics
	//get all interfaces
	interfaces, _ := net.Interfaces()
	for _, i := range interfaces {
		interfaceStatus, _ := net.IOCountersByFile(true, i.Name)
		for _, iStauts := range interfaceStatus {
			// in case of collecting duplicated interfaces
			if tmpSet.Contains(iStauts.Name) {
				continue
			}
			metric := NetMetrics{
				Name:        iStauts.Name,
				BytesSent:   float64(iStauts.BytesSent),
				BytesRecv:   float64(iStauts.BytesRecv),
				PacketsSent: float64(iStauts.PacketsSent),
				PacketsRecv: float64(iStauts.PacketsRecv),
				Errin:       float64(iStauts.Errin),
				Errout:      float64(iStauts.Errout),
				Dropin:      float64(iStauts.Dropin),
				Dropout:     float64(iStauts.Dropout),
				Fifoin:      float64(iStauts.Fifoin),
				Fifoout:     float64(iStauts.Fifoout),
			}
			netMetrics = append(netMetrics, metric)
			tmpSet.Add(iStauts.Name)
		}

	}
	return netMetrics
}