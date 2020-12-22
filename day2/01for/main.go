package main

import "fmt"

func main() {

	//流程控制之跳出for循环
	for i:=0;i<10;i++{
		fmt.Println(i)
		if i==5 {
			break
		}
	}
	fmt.Printf("over")

	//当i=5事跳过此次for循环，继续下一次循环
	for i:=0;i<10;i++{
		if i==5 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Printf("over")
}
