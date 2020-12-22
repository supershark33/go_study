package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	x int64
	wg sync.WaitGroup
	lock sync.Mutex
)

func main() {
	wg.Add(100000)
	for i:=0;i<100000;i++{
		go add()
	}
	wg.Wait()
	fmt.Println(x)
}



func add() {
	//单个的++不是线程安全的
	//x++;

	//用互斥锁解决问题
	//lock.Lock()
	//x++
	//lock.Unlock()

	//用atomic包解决问题
	atomic.AddInt64(&x, 1)
	wg.Done()
}
