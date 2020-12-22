package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLogger struct {
	level *logLevel
	path string
	fileName string
	maxFileSize int64
	fileObj *os.File
	errFileObj *os.File
}

//构造函数
func NewFileLogger(levelStr, path, fileName string, maxFileSize int64) (*FileLogger) {

	logLevel, err := getLogLevelByDesc(levelStr)
	if err != nil {
		panic(err)
	}
	flog := &FileLogger{
		level:       logLevel,
		path:        path,
		fileName:    fileName,
		maxFileSize: maxFileSize,
	}
	err = flog.initFile()
	if err != nil {
		flog.close()
		panic(err)
	}
	return flog
}

func (flog *FileLogger)Debug(format string, a... interface{}) {
	flog.doLog(DEBUG, format, a...)
}

func (flog *FileLogger)Info(format string, a... interface{}) {
	flog.doLog(INFO, format, a...)
}

func (flog *FileLogger)Warn(format string, a... interface{}) {
	flog.doLog(WARN, format, a...)
}

func (flog *FileLogger)Error(format string, a... interface{}) {
	flog.doLog(ERROR, format, a...)
}

func (flog *FileLogger)Fatal(format string, a... interface{}) {
	flog.doLog(FATAL, format, a...)
}

func (flog *FileLogger)checkSize(file *os.File) (*os.File, bool){
	stat, err := file.Stat()
	if err != nil || flog.maxFileSize > stat.Size(){
		return nil, false
	}
	//需要进行切割
	logName := path.Join(flog.path, stat.Name())
	bakLogName := fmt.Sprintf("%s.bak%s", logName, time.Now().Format("20060102150405000"))
	os.Rename(logName, bakLogName)
	//打开新的文件
	file, err = os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil{
		return nil, false
	}
	return file, true
}



func (flog *FileLogger)doLog(targetLevel logLevel, format string, a ...interface{}) {

	if !flog.level.enable(targetLevel){
		return
	}
	time := time.Now().Format("2006-01-02 15:04:05.000")
	msg := fmt.Sprintf(format, a...)
	funcName, fileName, lineNum := getInfo(3)
	if fileObj,ok := flog.checkSize(flog.fileObj); ok{
		flog.fileObj = fileObj
	}
	fmt.Fprintf(flog.fileObj, "[%s][%s][%s:%s():%d]%s\n", time, targetLevel.Desc, fileName, funcName, lineNum, msg)
	if targetLevel.Code < ERROR.Code {
		return
	}
	if fileObj,ok := flog.checkSize(flog.errFileObj); ok{
		flog.errFileObj = fileObj
	}
	fmt.Fprintf(flog.errFileObj, "[%s][%s][%s:%s():%d]%s\n", time, targetLevel.Desc, fileName, funcName, lineNum, msg)
}

func (flog *FileLogger) initFile() (err error) {
	fullName := path.Join(flog.path, flog.fileName)
	flog.fileObj, err = os.OpenFile(fullName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	flog.errFileObj, err = os.OpenFile(fullName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	//日志文件都已经打开
	return nil
}

func (flog FileLogger)close() {
	flog.fileObj.Close()
	flog.errFileObj.Close()
}







