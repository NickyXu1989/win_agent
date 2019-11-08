package routers

import (
	"github.com/gin-gonic/gin"
	"monkey_win_agent/controllers"
)

var metricsController controllers.MetricsController

//func NewTestRouters(r *gin.Engine) {
//	test := r.Group("/test")
//	{
//		test.GET("", controllers.TestController)
//	}
//
//}

func NewMetricsRouters(r *gin.Engine) {
	metrics := r.Group("/agent/metrics")
	{
		metrics.GET("/base", metricsController.BaseMetrics)
		metrics.GET("/cpu", metricsController.CpuMetrics)
		metrics.GET("/memory", metricsController.MemoryMetrics)
		metrics.GET("/diskUsage", metricsController.DiskMetrics)
		metrics.GET("/diskIO", metricsController.DiskIOMetrics)
		metrics.GET("/process", metricsController.ProcessMetrics)
		metrics.GET("/net", metricsController.NetMetrics)
	}
}
