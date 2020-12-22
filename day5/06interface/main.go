package main

import "fmt"

//interface 值接收者和指针接收者
//使用值接收者实现接口与使用指针接收者实现接口的区别?
//使用值接收者实现接口，结构体类型和结构体类型指针类型的变量都能存
//使用资金接收者实现的接口只能存结构体资质恩类型的变量

//接口和类型的关系
//同一个结构体可以实现多个接口
//接口还可以嵌套

//关于接口要注意的是，只有当有两个或者以上的具体类型必须以相同的方式进行处理或者约束的时候才需要定义接口
//不要为了定义接口而写接口，那样只会增加不必要的抽象，导致不必要的运行是的损耗

type animal interface {
	move()
	eat()
}

type cat struct {
	name string
	feet int8
}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("鸡动")
}

func (c chicken) eat() {
	fmt.Println("吃鸡饲料")
}

func (c cat)move()  {
	fmt.Println("猫动")
}

func (c cat)eat()  {
	fmt.Println("猫吃猫鼠")
}

func main() {

	var a animal
	c := cat{
		name: "猫咪",
		feet: 4,
	}
	a = c
	a.move()
	a.eat()
}