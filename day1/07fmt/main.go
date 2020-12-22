package main

import "fmt"

//fmt 占位符
func main() {
	var n = 100
	fmt.Printf("%T\n", n)
	fmt.Printf("%v\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	s1 := "hello 沙河"
	fmt.Printf("%T\n", s1)
	fmt.Printf("%s\n", s1)
	fmt.Printf("%v\n", s1)
}
