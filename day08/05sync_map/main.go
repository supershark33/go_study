package main

import (
	"fmt"
	"strconv"
	"sync"
)

var (
	holder = make(map[string]int)
	safeHolder = sync.Map{}
)

//go内置的map不是并发安全的，会产生concurrent map writes的问题
//解决方案
//1、互斥锁 2、读写锁 3、sync.Map
//像这种场景下就需要为map加锁来保证并发的安全性了，Go语言的sync
//包中提供了一个开箱即用的并发安全版map-sync .Map 。 ，开箱即用表
//示不用像内置的map-样使用make函数初始化就能直接使用。
//同时sync .Map内置了诸如Store 、 Load| 、 LoadOrStore|、 Delete Range等操作方法。



//func main() {
//	wg := sync.WaitGroup{}
//	for i:=0;i<3;i++{
//		wg.Add(1)
//		go func(n int) {
//			key := strconv.Itoa(n) //将int转化为string
//			set(key, n)
//			fmt.Printf("k=%v,v=%v\n",key, get(key))
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//}

func main() {
	wg:= sync.WaitGroup{}
	for i:=0;i<50;i++{
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			safeHolder.Store(key, n) //必须用sync.map的store方法取值
			value, _ := safeHolder.Load(key) //必须用sync.map的load方法取值
			fmt.Printf("k=%v,v=%v\n",key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}


func set(key string, value int){
	holder[key] = value
}

func get(key string) int {
	return holder[key]
}
