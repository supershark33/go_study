package mylogger

import (
	"errors"
	"fmt"
)

type logLevel struct {
	Code uint8
	Desc string
}

var (
	DEBUG = logLevel{
		Code: 3,
		Desc:"DEBUG",
	}
	TRACE = logLevel{
		Code: 6,
		Desc:"TRACE",
	}
	INFO = logLevel{
		Code: 9,
		Desc:"INFO",
	}
	WARN = logLevel{
		Code: 12,
		Desc:"WARN",
	}
	ERROR = logLevel{
		Code: 15,
		Desc:"ERROR",
	}
	FATAL = logLevel{
		Code: 18,
		Desc:"FATAL",
	}
	CDMAP = map[string]*logLevel{
		DEBUG.Desc:&DEBUG,
		TRACE.Desc:&TRACE,
		INFO.Desc:&INFO,
		WARN.Desc:&WARN,
		ERROR.Desc:&ERROR,
		FATAL.Desc:&FATAL,
	}
)

func getLogLevelByDesc(logDesc string) (level *logLevel, err error) {

	if value, ok := CDMAP[logDesc]; ok {
		return value, nil
	}else {
		return &INFO, errors.New(fmt.Sprintf("unkown loglevel : %s, default level INFO seted, loglevel should be one of (DEBUG,TRACE,INFO,WARN,ERROR,FATAL)", logDesc))
	}
}

//enable 判断日志是否打印
func (format logLevel)enable(target logLevel) bool{
	return target.Code >= format.Code
}
