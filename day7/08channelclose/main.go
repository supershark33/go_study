package main

import "fmt"

//关闭通道,关闭通道将不能向通道内部写入，只能向通道内部读取

func main() {

	cha := make(chan int, 2)
	cha <- 10
	cha <- 20
	close(cha) //如果不close, 那么后续的for 和 for range 将会产生异常
	//for x := range cha {
	//	fmt.Println(x)
	//}

	y, ok := <- cha
	fmt.Println(y, ok)
	<- cha
	<- cha //不会产生异常，就算取完了
	x, ok := <- cha
	fmt.Println(x, ok) //依然可以取值
}
