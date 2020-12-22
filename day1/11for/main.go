package main

import "fmt"

//golang中只有一个for循环，但是变种非常多
func main() {

	//基本格式
	for i1:=0;i1<10;i1++{
		fmt.Println(i1)
	}

	//变种一
	i2 := 5
	for ; i2 <10; i2++{
		fmt.Println(i2)
	}
	
	//变种三
	i3 := 5
	for i3<10{
		fmt.Println(i3)
		i3++
	}

	//变种四
	//for{
	//	fmt.Println("123")
	//}

	//for range
	s := "hello刘云克"
	for i,v := range s{
		fmt.Printf("%d %c\n",i,v)
	}
}
