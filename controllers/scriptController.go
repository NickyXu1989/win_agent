package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"monkey_win_agent/utils"
	"net/http"
)


type RunScriptReq struct {
	MonId string `json:"monId"`
	ScriptName string `json:"scriptName"`
	ScriptText string `json:"scriptText"`
	ScriptMD5 string `json:"md5"`
	FireCmd string `json:"fireCmd"`
	Timeout int `json:"timeout"`
}


type ScriptController struct {

}


// push script and run
func (ctl *ScriptController)RunScript(c *gin.Context){
	//monId := c.PostForm("monId")
	//scriptName := c.PostForm("scriptName")
	//scriptText := c.PostForm("scriptText")
	//scriptMD5 := c.PostForm("md5")
	//fireCmd := c.PostForm("fireCmd")
	//timeoutStr := c.PostForm("timeout")
	var runScriptReq RunScriptReq
	c.BindJSON(&runScriptReq)

	fmt.Println(runScriptReq.MonId)
	fmt.Println(runScriptReq.ScriptName)
	fmt.Println(runScriptReq.ScriptText)
	fmt.Println(runScriptReq.ScriptMD5)
	success := true
	errors := ""
	result := ""

	scriptHelper := utils.NewScriptFileHelperInstance()
	if success {
		err := scriptHelper.AddOrUpdateScriptFile(runScriptReq.MonId, runScriptReq.ScriptName, runScriptReq.ScriptText, runScriptReq.ScriptMD5)
		if err != nil {
			success = false
			errors = errors + err.Error()
		}
	}


	if success {
		scriptResult, err := scriptHelper.RunScript(runScriptReq.MonId,runScriptReq.ScriptName, runScriptReq.ScriptMD5, runScriptReq.FireCmd, runScriptReq.Timeout)
		if err != nil {
			success = false
			errors = errors + err.Error()
		}
		result = scriptResult
	}
	c.JSON(http.StatusOK, gin.H{"success": success, "result": result, "errorMsg": errors})

}


//use local cache and monId to run script
func (ctl *ScriptController)RunScriptWithLocalMonId(c *gin.Context){
	//monId := c.PostForm("monId")
	//scriptMD5 := c.PostForm("md5")
	//fireCmd := c.PostForm("fireCmd")
	//timeoutStr := c.PostForm("timeout")
	//
	//timeout, timeoutError := strconv.Atoi(timeoutStr)
	//if timeoutError != nil {
	//	panic(timeoutError.Error())
	//}
	var runScriptReq RunScriptReq
	c.BindJSON(&runScriptReq)
	scriptHelper := utils.NewScriptFileHelperInstance()
	success := true
	result, err := scriptHelper.RunScript(runScriptReq.MonId, runScriptReq.ScriptName,runScriptReq.ScriptMD5, runScriptReq.FireCmd, runScriptReq.Timeout)
	if err != nil {
		success = false
		c.JSON(http.StatusOK, gin.H{"success":success, "result":result, "errorMsg": err.Error()})
	}else {
		c.JSON(http.StatusOK, gin.H{"success": success, "result": result})
	}
}
