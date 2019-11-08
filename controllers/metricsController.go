package controllers

import (
	"github.com/gin-gonic/gin"
	"monkey_win_agent/utils"
	"net/http"
	"os"
)

type MetricsController struct {

}

func (ctl *MetricsController)BaseMetrics(c *gin.Context) {
	result := map[string]interface{}{}
	var metrics []map[string]interface{}
	//get hostname
	hostname, _ := os.Hostname()
	result["host"] = hostname

	//get cpu percent
	cpuInstance := utils.GetCpuMetricsHistoryInstance()
	cpuPct := cpuInstance.GetLatestCpuPercent()
	cpuMetric := map[string]interface{}{
		"metric": "cpu",
		"value": cpuPct,
	}
	metrics = append(metrics, cpuMetric)


	//get memory percent
	memPct := utils.GetMemPct()
	memMetric := map[string]interface{}{
		"metric": "memory",
		"value" : memPct,
	}
	metrics = append(metrics, memMetric)

	//get disk usage
	diskMetrics := map[string]interface{}{
		"metric": "disk",
		"data": utils.GetDiskUsagePct(),
	}
	metrics = append(metrics, diskMetrics)

	result["metrics"] = metrics
	c.JSON(http.StatusOK, result)
}


//cpu metrics
func (ctl *MetricsController)CpuMetrics(c *gin.Context) {
	result := map[string]interface{}{}
	var metrics []map[string]interface{}
	//get hostname
	hostname, _ := os.Hostname()
	result["host"] = hostname

	//get cpu percent
	cpuInstance := utils.GetCpuMetricsHistoryInstance()
	cpuPct := cpuInstance.GetLatestCpuPercent()
	cpuMetric := map[string]interface{}{
		"metric": "cpu",
		"value": cpuPct,
	}
	metrics = append(metrics, cpuMetric)
	result["metrics"] = metrics
	c.JSON(http.StatusOK, result)
}


//memory metrics
func (ctl *MetricsController)MemoryMetrics(c *gin.Context) {
	result := map[string]interface{}{}
	var metrics []map[string]interface{}
	//get hostname
	hostname, _ := os.Hostname()
	result["host"] = hostname

	//get memory percent
	memPct := utils.GetMemPct()
	memMetric := map[string]interface{}{
		"metric": "memory",
		"value" : memPct,
	}
	metrics = append(metrics, memMetric)
	result["metrics"] = metrics
	c.JSON(http.StatusOK, result)
}


//disk usage
func (ctl *MetricsController)DiskMetrics(c *gin.Context) {
	result := map[string]interface{}{}
	var metrics []map[string]interface{}
	//get hostname
	hostname, _ := os.Hostname()
	result["host"] = hostname

	//get disk usage
	diskMetrics := map[string]interface{}{
		"metric": "disk",
		"data": utils.GetDiskUsagePct(),
	}
	metrics = append(metrics, diskMetrics)

	result["metrics"] = metrics
	c.JSON(http.StatusOK, result)
}



//disk io metrics
func (ctl *MetricsController)DiskIOMetrics(c *gin.Context) {
	result := map[string]interface{}{}
	var metrics []map[string]interface{}
	//get hostname
	hostname, _ := os.Hostname()
	result["host"] = hostname

	//get disk metrics
	diskIOMetrics := map[string]interface{}{
		"metric": "diskIO",
		"data": utils.GetDiskIOMetrics(),
	}

	metrics = append(metrics, diskIOMetrics)

	result["metrics"] = metrics
	c.JSON(http.StatusOK, result)
}


//net metrics
func (ctl *MetricsController)NetMetrics(c *gin.Context) {
	result := map[string]interface{}{}
	var metrics []map[string]interface{}
	hostname,_ := os.Hostname()
	result["host"] = hostname

	//get net process
	netMetrics := map[string]interface{}{
		"metric": "net",
		"data": utils.GetNetStatus(),
	}

	metrics = append(metrics, netMetrics)

	result["metrics"] = metrics
	c.JSON(http.StatusOK, result)
}

//process metrics
func (ctl *MetricsController)ProcessMetrics(c *gin.Context) {
	result := map[string]interface{}{}
	var metrics []map[string]interface{}
	hostname,_ := os.Hostname()
	result["host"] = hostname

	//get process process
	processMetrics := map[string]interface{}{
		"metric": "process",
		"data": utils.GetProcessStatus(),
	}

	metrics = append(metrics, processMetrics)

	result["metrics"] = metrics
	c.JSON(http.StatusOK, result)
}


