package main

import (
	"fmt"
	"sync"
	"time"
)

//线程间的第二种交互方式，通过channel传递信号

var (
	wg sync.WaitGroup
	exitChan chan bool = make(chan bool, 1)
)


func main() {

	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	exitChan <- true
	wg.Wait()
}

func f() {
	defer wg.Done()
	for {
		count := 1
		fmt.Println("张拾玥，", count)
		count++
		time.Sleep(time.Millisecond * 500)
		select {
		case <-exitChan:
			return
		default:
		}
	}
}
