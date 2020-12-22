package main

import "fmt"

//常量
//定义常量之后将不能修改
//在程序运行期间不会改变的量
const pi = 3.1415926

//常量批量生平
const (
	statusOk = 200
	notFound = 404
)

//批量声明常量后，后面的值默认取全面的值
const (
	n1 = 100
	n2
	n3
)

//iota
const (
	a1 = iota  //0
	a2         //1
	a3		   //2
)

//iota例子1,多个iota和_变量占位符
const (
	b1 = iota  //0
	b2 = iota  //1
	_
	b3 = iota  //3
)

const (
	d1,d2 = iota + 1, iota + 2  //d1 = 1 d2 = 2
	d3,d4 = iota + 1, iota + 2  //d3 = 2 d4 = 3
)

const (
	_ = iota //第一个直接扔了
	KB = 1 << (10 * iota)  //1左移十位，相当于2的十次方
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)


func main() {
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)

	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)
	fmt.Println(d4)
}
