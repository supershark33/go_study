package main

import (
	"fmt"
	"sync"
)

func main() {
	//noBufChannel()
	bufChannel()
}

var wg sync.WaitGroup

func bufChannel() {

	//定义channel
	var ch2 chan int
	ch2 = make(chan int, 2)
	b := 2
	ch2 <- b //存放成功
	ch2 <- 3 //存放成功
	//ch2 <- 4 //存放失败，all goroutines are asleep - deadlock!
	fmt.Println("数据存放成功", b)
	close(ch2)
}

func noBufChannel() {

	//定义channel
	var ch1 chan int //nil
	ch1 = make(chan int) //不带缓冲区的初始化
	//ch1 <- 1 // all goroutines are asleep - deadlock! 不带缓存区的直接放会有问题
	//x := <-ch1
	b := 10
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <- ch1
		fmt.Printf("参数%d已经被接受\n", x)
	}()
	ch1 <- b
	fmt.Printf("参数%d已经被发送\n", b)
	wg.Wait()
}
