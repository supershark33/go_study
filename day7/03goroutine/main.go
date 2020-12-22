package main

import (
	"fmt"
	"time"
)

//程序启动的时候，main函数也是一个goroutine，main的goroutine
//并发与并行
//并发:同一时间段内执行多个任务(你在用微信和两个女朋友聊天)。
//并行:同一时刻执行多个任务(你和你朋友都在用微信和女朋友聊天)。
//Go语言的并发通过goroutine| 实现。goroutine 类似于线程，属于用户态的线程，我们可以根据需要创建成千上万个| goroutine并发工作。goroutine 是由Go语言的运行时(runtime) 调度完成，而线程是由操作系统调度完成。Go语言还提供channel| 在多个| goroutine|间进行通信。| goroutine |和channel| 是Go语言秉承的
//CSP (Communicating Sequential Process)并发模式的重要实现基础。
//goroutine
//在java/c+ +中我们要实现并发编程的时候，我们通常需要自己维护-一个线程池，并且需要自己去包装一个又- 一个的任务,同时需要自己去调度线程执行任务并维护上下文切换，这一切通常会耗费程序员大量的心智。那么能不能有一种机制，程序员只需要定义很多个任务，让系统去帮助我们把这些任务分配到CPU上实现并发执行呢?
//Go语言中的| goroutine |就是这样-种机制，| goroutine |的概念类似于线程，但goroutine| 是由Go的运行时
//(runtime)调度和管理的。Go程序会智能地将goroutine中的任务合理地分配给每个CPU。Go语言之所以被称为现代化的编程语言，就是因为它在语言层面已经内置了调度和上下文切换的机制。
//在Go语言编程中你不需要去自己写进程、线程、协程，你的技能包里只有一一个技能- goroutine| ，当你需要让某个任务并发执行的时候，你只需要把这个任务包装成一个函数，开启-一个| goroutine|去执行这个函数就可以了， 就是这么简单粗暴。
//使用goroutine
//Go语言中使用| goroutine非常简单，只需要在调用函数的时候在前面加上go关键字，就可以为一一个函数创建一个goroutine。
//一个| goroutine |必定对应-一个函数，可以创建多个| goroutine去执行相同的函数。
//
//goroutine与线程
//可增长的栈
//OS线程(操作系统线程) -般都有固定的栈内存(通 常为2MB) ,- 个goroutine| 的栈在其生命周期开始时只有很小的栈(典型情况下2KB)，goroutine 的栈不是固定的，他可以按需增大和缩小, goroutine |的栈大小限制可以达到1GB，虽然极少会用到这个大。所以在Go语言中- -次创建十万左右的 goroutine |也是可以的。
//goroutine调度
//GMP |是Go语言运行时(runtime) 层面的实现，是go语言自己实现的一套调度系统。区别于操作系统调度OS线程。
//● G很好理解，就是个goroutine的， 里面除了存放本goroutine信息外还有与所在P的绑定等信息。
//M (machine)是Go运行时(runtime) 对操作系统内核线程的虚拟，M与内核线程一 般是一映射的关系,- -个groutine最终是要放到M.上执行的;
//● P管理着-组goroutine队列， P里面会存储当前goroutine运行的上下文环境(函数指针,堆栈地址及地址边
//界)，P会对自己管理的goroutine队列做一些调度(此如把占用CPU时间较长的goroutine暂停、运行后续的goroutine等等)当自己的队列消费完了就去全局队列里取，如果全局队列里也消费完了会去其他P的队列里抢任务。
//P与M-般也是一-对应的。 他们关系是: P管理着-组G挂载在M 上运行。当一个G长久阻塞在一个M上时，runtime会新建一个M，阻塞G所在的P会把其他的G挂载在新建的M.上。当旧的G阻塞完成或者认为其已经死掉时回收旧的M。
//
//Go语言中的操作系统线程和goroutine的关系:
//1. -一个操作系统线程对应用户态多个gorouting。2. go程序可以同时使用多个操作系统线程。
//3. goroutine和OS线程是多对多的关系，即m:n。
//
//gorouti ne什么结束?
//goroutine对应的函数结束了，goroutine结束 了。
//main函数执行完了，由main函数创建的那些gor outine都结束了。


func main() {
	for i:=0;i<100;i++{
		go hello(i) //开启一个单独的goroutine去执行这个函数
	}
	fmt.Println("main")
	//main函数退出，其线下面的goroutinne就都退出了
	time.Sleep(time.Second)
}

func hello(i int) {
	fmt.Println("你好，张拾玥", i)
}


