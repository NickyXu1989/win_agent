package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"monkey_win_agent/routers"
	"monkey_win_agent/utils"
	_ "net/http/pprof"
	"time"
)

type Server struct {
	router *gin.Engine
}

var server *Server
var cpuMetricsHistory *utils.CpuMetricsHistory

var ip string
var port string

func main() {

	//pprof
	//go func() {
	//	http.ListenAndServe("0.0.0.0:8899", nil)
	//}()

	ip := flag.String("ip", "0.0.0.0", "http listen ip")
	port := flag.String("port", "7878", "http listen port")
	flag.Parse()

	listenAddr := *ip + ":" + *port


	scriptFileHelper := utils.NewScriptFileHelperInstance()
	scriptFileHelper.InitScriptFiles()

	cpuMetricsHistory = utils.NewCpuMetricsHistory()

	cpuMetricsHistory.GatherCpuMetric()
	//timer loop for cpu metrics
	ticker := time.NewTicker(time.Second * 15)

	go func() {
		for {
			select {
			case <- ticker.C:
				cpuMetricsHistory.GatherCpuMetric()
			}
		}
	}()

	defer func() {
		if p := recover();p != nil {
			log.Fatal(p)
		}
	}()

	server := Server{}
	server.router = gin.Default()

	//routers.NewTestRouters(server.router)
	routers.NewMetricsRouters(server.router)
	routers.NewScriptRouters(server.router)
	server.router.Run(listenAddr)

}

