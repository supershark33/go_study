package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Printf("work: %d, start job: %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("work: %d, finish job: %d\n", id, j)
		result <- j * 2
		wg.Done()
	}
}

var wg sync.WaitGroup

func main() {

	jobs := make(chan int, 100)
	results := make(chan int, 100)
	//开启3个gorouting,也就是线程池里面有三个线程
	for i:=0;i<3;i++{
		go worker(i, jobs, results)
	}
	//分配五个任务到线程池
	for i:=0;i<100;i++{
		wg.Add(1)
		jobs <- i
	}
	//执行完成后关闭线程池
	close(jobs)
	wg.Wait()
	//for a:= 1; a<=5; a++{
	//	<- results
	//}
	//time.Sleep(time.Second * time.Duration(10))

}
