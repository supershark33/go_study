package main

import "fmt"

func main() {
	//math.MaxFloat32
	f1 := 1.23456
	fmt.Println(f1)
	f2 := float32(1.23456)
	fmt.Println(f2)
	//f1 = f2 32和64位的浮点数，两者完全不同，不能互相赋值
}

