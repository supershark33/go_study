package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	notify bool
	wg sync.WaitGroup
)

//线程间的第一种通信方式，全局变量

func main() {

	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	notify = true
	wg.Wait()
}

func f() {

	defer wg.Done()
	count := 1
	for {
		fmt.Println("张拾玥，", count)
		count++
		time.Sleep(time.Millisecond * 500)
		if notify {
			break
		}

	}
}
