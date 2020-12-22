package main

import (
	"fmt"
	"time"
)

func main() {

	//获取当前时间的时分秒
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	//获取当前时间戳
	unix := now.Unix()
	fmt.Println(unix)
	nano := now.UnixNano()
	fmt.Println(nano)

	//当前时间戳转换为日期
	t := time.Unix(1606356165, 0)
	fmt.Println(t)
	fmt.Println(t.Year())
	fmt.Println(t.Month())
	fmt.Println(t.Day())

	//时间间隔
	fmt.Println(time.Second)
	fmt.Println(time.Hour)
	fmt.Println(time.Microsecond)

	//时间操作
	fmt.Println(now.Add(24*time.Hour))

	//定时器
	//timer := time.Tick(time.Second)
	//for t := range timer{
	//	fmt.Println(t)
	//}

	//时间格式化
	//时间类型有一一个自带的方法Format| 进行格式化，需要注意的是Go语言中格式化时间模板不是常见的|
	//Y-m-d H:M:S而是使用Go的诞生时间2006年1月2号15点o4分(记忆口诀为20061234) 。也许这就是技
	//术人员的浪漫吧。 补充:如果想格式化为12小时方式，需指定| PM|。
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("2006/01/02"))
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006/01/02 03:04:05 PM"))
	fmt.Println(now.Format("2006/01/02 15:04:05.000"))

	//将按照对应对应的格式解析时间
	parse, err := time.Parse("2006-01-02", "2000-08-03")
	if(err != nil){
		fmt.Println("格式化时间错误")
		return
	}
	fmt.Println(parse)
	fmt.Println(parse.Unix())

	//求两个时间的差值
	sub := now.Sub(now.Add(24 * time.Hour))
	fmt.Println(sub)

	//sleep
	//n := 5
	//time.Sleep(time.Duration(n) * time.Second)
	//fmt.Println("让出cpu了5s")
	//time.Sleep(5 * time.Second)
	//fmt.Println("让出cpu了5s")

	//关于时区
	t2 := time.Now() //获取的是当前时区的时间，即东八区
	fmt.Println(t2)
	//parse的，却是UGC的时间，不是东八区
	time.Parse("2006-01-02 15:04-05", "2020-11-26 14:37:46") //UGC的时间
	//按照东八区的时区和格式去解析字符串的实际时间
	//先指定时区
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("load loc failed: %v\n", err)
		return
	}
	inLocation, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-11-26 14:37:46", location)
	if err != nil {
		fmt.Printf("parse time error: %v\n", inLocation)
		return
	}
	//计算差值
	fmt.Println(now.Sub(inLocation))
}
