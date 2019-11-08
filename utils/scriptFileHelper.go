package utils

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"
)

type ScriptFileHelper struct {
	fileRoot string
}

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)



var lock sync.Mutex
var scriptFileHelperInstance *ScriptFileHelper

func NewScriptFileHelperInstance() *ScriptFileHelper {
	if scriptFileHelperInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		scriptFileHelperInstance = &ScriptFileHelper{
			fileRoot: os.TempDir() + "\\" + "win_monkey_agent",
		}
	}
	return scriptFileHelperInstance
}

func GetScriptFileHelperInstance () *ScriptFileHelper {
	if scriptFileHelperInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		scriptFileHelperInstance = &ScriptFileHelper{
			fileRoot: os.TempDir() + "\\" + "win_monkey_agent",
		}
	}
	return scriptFileHelperInstance
}


func (c *ScriptFileHelper)InitScriptFiles() {
	file, _ := os.Stat(c.fileRoot)
	if file == nil {
		err := os.Mkdir(c.fileRoot, os.ModePerm)
		if err != nil {
			panic(err.Error())
		}
	}

}

func (c *ScriptFileHelper)AddOrUpdateScriptFile(monId string ,fileName string, fileContent string,fileHash string) error{

	//first, check the hash and fileContent
	file, err := ioutil.TempFile("",  fileName+"*"+".bat")
	if err != nil {
		panic(err.Error())
	}
	defer os.Remove(file.Name())

	//add scriptText to the file
	fmt.Println(file.Name())
	ioutil.WriteFile(file.Name(),[]byte(fileContent),os.ModePerm)
	file.Close()


	//get the md5
	data,_ := ioutil.ReadFile(file.Name())
	md5str := GetMD5Encode(string(data[:]))
	//fmt.Println(md5str)





	if md5str != fileHash {
		fmt.Println("first hash check error")
		return errors.New("fileHash and fileContent not compatible")
	}


	dirName := c.fileRoot + "\\" + monId
	filePath := c.fileRoot + "\\" + monId + "\\" + fileName

	//recreate the file
	dir, _ := os.Stat(dirName)
	if dir != nil {
		os.RemoveAll(dirName)
		err := os.Mkdir(dirName, os.ModePerm)
		if err != nil {
			panic(err)
		}
	} else {
		os.Mkdir(dirName, os.ModePerm)
	}
	//rewrite the script
	_ = ioutil.WriteFile(filePath, []byte(fileContent), os.ModePerm)

	return nil
}

func (c *ScriptFileHelper)RunScript(monId string, scriptName string, scriptHash string, scriptType string, timeout int) (result string, err error) {
	instance := GetScriptFileHelperInstance()
	scriptPath := instance.fileRoot + "\\" + monId + "\\" + scriptName
	//fmt.Println(scriptPath)
	_, err = os.Stat(scriptPath)
	if err != nil {
		return "", errors.New("no script")
	}
	switch scriptType {
	case "cmd":
		if !checkScriptHash(scriptPath, scriptHash) {
			return "", errors.New("hash error")
		}
		result, err := runBatScript(scriptPath,timeout)
		if err != nil {
			return result, err
		}
		return result, nil
	}
	return "", errors.New("type not supported")

}


func checkScriptHash(scriptPath string, scriptHash string) bool {
	data,_ := ioutil.ReadFile(scriptPath)
	md5Str := GetMD5Encode(string(data))
	//fmt.Println()
	//fmt.Println(md5Str)
	if md5Str != scriptHash {
		return false
	}
	return true
}


//run cmd script
func runBatScript(scriptPath string, timeout int) (scriptOutput string, err error) {
	//ready to run the script
	result := ""

	//run the script
	cmd := exec.Command("cmd.exe", "/c ",scriptPath)

	//define the output pipe and err pipe
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println(err)
		return
	}

	cmd.Start()

	done := make(chan error)
	go func() {
		//convert strings from GBK to UTF-8
		outbuf := bufio.NewScanner(stdout)
		for outbuf.Scan() {
			cmdRe:=convertByte2String(outbuf.Bytes(),"GB18030")
			result = result + cmdRe + "\n"
		}

		errbuf := bufio.NewScanner(stderr)
		for errbuf.Scan() {
			cmdRe:=convertByte2String(errbuf.Bytes(),"GB18030")
			result = result + cmdRe  + "\n"
		}
		done <- cmd.Wait()
	}()
	// Start a timer
	timeoutEvent := time.After(time.Duration(timeout) * time.Second)

	// The select statement allows us to execute based on which channel
	// we get a message from first.
	select {
	case <-timeoutEvent:
		// Timeout happened first, kill the process and print a message.
		cmd.Process.Kill()
		return "", errors.New("timeout")
	case err := <-done:
		return result, err
	}
}


func scanDir(dirName string) []string {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Println(err)
	}

	var fileList []string
	for _, file := range files {
		fileList = append(fileList, dirName + string(os.PathSeparator) + file.Name())
	}
	return fileList
}



func scanDirs(dirName string) []string {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Println(err)
	}
	var fileList []string
	for _, file := range files {
		fileList = append(fileList, dirName + string(os.PathSeparator) + file.Name())
		if file.IsDir() {
			fileList = append(fileList, scanDir(dirName + string(os.PathSeparator) + file.Name())...)
		}
	}
	return fileList
}



func convertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes,_=simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str= string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}


//return a 32byte md5 string
func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

//return a 16byte md5 string
func Get16MD5Encode(data string) string{
	return GetMD5Encode(data)[8:24]
}
