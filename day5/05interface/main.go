package main

import "fmt"

//接口是一种类型

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
