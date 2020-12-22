package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	//random()
	for i:=0;i<10;i++{
		wg.Add(1)
		go f1(i)
	}
	wg.Wait()
}

func random() {

	//随机数种子是为了让每次产生的随机数不一样
	rand.Seed(time.Now().UnixNano())
	for i:= 0;i<10;i++{
		r1 := rand.Int() //int64的随机数
		r2 := rand.Intn(99) //一定范围的随机数
		fmt.Println(r1, r2)
	}
}

var wg sync.WaitGroup

func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	fmt.Println(i)
}
