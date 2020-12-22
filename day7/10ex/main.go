package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
使用goroutine和channel实现一个计算int64随机数名位数和的程序。
1.开启一个goroutine循环生 成int64类型的随机数，发送到jobChan
2.开启24个goroutine从jobchan中取出随机数计算各位数的和，将结果发送到resultChan
3.主goroutine从resultChan取出结果并打印到终端输出
 */
func main() {

	jobs := make(chan int64, 100)
	results := make(chan int64, 100)
	go randomNum(jobs)
	for i:=0;i<24;i++{
		go sumBits(jobs, results)
	}
	wg.Wait()
	for count := range results{
		fmt.Printf("得到的加和为：%d\n", count)
	}
}

var wg sync.WaitGroup

func randomNum(jobs chan<- int64) {
	rand.Seed(time.Now().UnixNano())
	wg.Add(5)
	for i:=0;i<5;i++{
		n := rand.Int63()
		fmt.Println(n)
		jobs <- n
	}
	close(jobs)
}

func sumBits(jobs <-chan int64, results chan<- int64) {
	for num := range jobs{
		results <- addAllbits(num)
		wg.Done()
	}
}

func addAllbits(num int64) (count int64) {
	for num > 0{
		count += num%10
		num = num/10
	}
	return
}