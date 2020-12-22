package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second*5)
	cancel()
	wg.Wait()
	fmt.Println("over")
}

func worker(ctx context.Context) {

	defer wg.Done()
	for {
		fmt.Println("db connenting...")
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done(): //50ms后自动调用
			return
		default:
		}
	}
	fmt.Println("worder done!")
}
