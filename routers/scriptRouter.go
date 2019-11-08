package routers

import (
"github.com/gin-gonic/gin"
"monkey_win_agent/controllers"
)

var scriptController controllers.ScriptController


func NewScriptRouters(r *gin.Engine) {
	metrics := r.Group("/agent")
	{
		metrics.POST("/script", scriptController.RunScript)
		metrics.POST("/monitor", scriptController.RunScriptWithLocalMonId)
	}
}
