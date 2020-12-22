package mylogger

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

func getInfo(skip int) (funcName, fileName string, lineNum int){
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime.Celler() failed")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]
	lineNum = line
	return
}


