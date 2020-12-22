package main

import (
	"log"
	"time"
)

//需求分析
//1、支持往不同的地方写
//2、支持不同的日志级别 debug trace info warning error fatal
//3、支持开关控制
//4、完整的日志记录要包含时间、行号、文件名、日志级别、日志信息
//5、日志文件要切割
func main()  {

	for{

		log.Println("这是一条测试的日志")
		time.Sleep(time.Second * 3)

	}
}
