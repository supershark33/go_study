package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//使用context做gorouting之间的通信，便于统一管理

var (
	wg sync.WaitGroup
)


func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f1(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
}

func f1(ctx context.Context) {
	defer wg.Done()
	go f2(ctx)
	count := 1
	for {
		fmt.Println("张拾玥，", count)
		count++
		time.Sleep(time.Millisecond * 500)
		select {
		case <- ctx.Done():
			return
		default:
		}
	}
}

func f2(ctx context.Context) {
	defer wg.Done()
	count := 1
	for {
		fmt.Println("小九九，", count)
		count++
		time.Sleep(time.Millisecond * 500)
		select {
		//等待上级的通知
		case <- ctx.Done():
			return
		default:
		}
	}
}
