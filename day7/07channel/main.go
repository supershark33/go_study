package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var once sync.Once

func f1(ch1 chan int) {
	defer wg.Done()
	for i:=0;i<100;i++{
		ch1 <- i
	}
	close(ch1) //如果不关闭ch1，那么下面对ch1的for range循环将会导致dead lock，可以用生硬的for 循环来避免
}

func f2(ch1,ch2 chan int)  {
	defer wg.Done()
	//非生硬的for循环，会由于ch1在数据没有close的情况下导致deadlock
	for x := range ch1{
		ch2 <- x*x
	}
	//生硬的for循环，可以加个判断，判断是否close
	//for {
	//	x, ok := <- ch1
	//	if !ok {
	//		break
	//	}
	//	ch2 <- x * x
	//}
	once.Do(func() {
		close(ch2)
	})
	//close(ch2) //确保某个操作只能做一次，如果做两次将会导致异常panic: close of closed channel
}

func main() {
	var ch1,ch2 chan int
	ch1 = make(chan int, 100)
	ch2 = make(chan int, 100)
	wg.Add(3)
	go f1(ch1)
	go f2(ch1, ch2)
	go f2(ch1, ch2)
	wg.Wait()
	for x := range ch2{
		fmt.Println(x)
	}
}


