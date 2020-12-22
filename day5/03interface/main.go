package main

import "fmt"

//接口也是一种类型，他约束了变量需要具备哪些方法

//定义接口
type speaker interface {
	speak() //只要实现了speak方法的struct就是speaker类型，接口是一个类型
}

type cat struct {}

type dog struct {}

type person struct {}

func (c cat)speak() {
	fmt.Println("喵喵喵~")
}

func (d dog)speak()  {
	fmt.Println("汪汪汪~")
}

func (p person)speak()  {
	fmt.Println("嘤嘤嘤~")
}

func da(speaker speaker)  {
	speaker.speak()
}

func main() {
	var c cat
	c.speak()
	da(c)

	var p person
	p.speak()
	da(p)

	var d dog
	d.speak()
	da(d)
}




