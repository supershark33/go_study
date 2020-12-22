package mylogger

import (
	"fmt"
	"time"
)

//构造函数
func NewLog(levelStr string) ConsoleLogger {
	return ConsoleLogger{}
}



//日志类
type ConsoleLogger struct {
	Level logLevel
}

func parseLogLevel()  {
	
}

func (log ConsoleLogger)doLog(targetLevel logLevel, format string, a... interface{}) {
	if log.Level.enable(targetLevel) {
		return
	}
	now := time.Now()
	funcName, fileName, num := getInfo(3)
	msg := fmt.Sprintf(format, a...)
	logBaseInfo := fmt.Sprintf("[%s][%s][%s:%s():%d]%s", now.Format("2006-01-02 15:04:05.000"), targetLevel.Desc, fileName, funcName, num, msg)
	fmt.Println(logBaseInfo)
}

func (log ConsoleLogger)Debug(format string, a... interface{}) {
	log.doLog(DEBUG, format, a...)
}

func (log ConsoleLogger)Info(format string, a... interface{}) {
	log.doLog(INFO, format, a...)
}

func (log ConsoleLogger)Warn(format string, a... interface{}) {
	log.doLog(WARN, format, a...)
}

func (log ConsoleLogger)Error(format string, a... interface{}) {
	log.doLog(ERROR, format, a...)
}

func (log ConsoleLogger)Fatal(format string, a... interface{}) {
	log.doLog(FATAL, format, a...)
}
