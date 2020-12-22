package main

import (
	"fmt"
	"zhangyuanbo.com/day6/mylogger"
)

func main() {

	logger := mylogger.NewFileLogger("INFO", "/Users/zhangyuanbo/Bytedance/src/zhangyuanbo.com/day6/mylogger/log","test.log", 10*1024*1024)

	for {
		logger.Debug("这是一条debug日志, name = %s, id = %d", "哈哈哈", 6421)
		logger.Info("这是一条info日志, name = %s, id = %d", "哈哈哈", 6421)
		logger.Warn("这是一条warn日志, name = %s, id = %d", "哈哈哈", 6421)
		logger.Error("这是一条error日志, name = %s, id = %d", "哈哈哈", 6421)
		logger.Fatal("这是一条fatal日志, name = %s, id = %d", "哈哈哈", 6421)
		fmt.Println("正在打印...")
		//time.Sleep(time.Second)
	}
}
