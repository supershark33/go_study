package main

import "fmt"

//构造函数

type person struct {
	name string
	age int
}

//构造函数
//返回结构体还是结构体指针，要看结构体的大小
func newPerson(name string, age int) person {
	return person{
		name : name,
		age: age,
	}
}

//返回结构体指针
func newPersonP(name string, age int) *person {
	return &person{
		name: name,
		age: age,
	}
}

func main() {
	p := newPerson("刘云克", 12)
	fmt.Println(p)
	personP := newPersonP("仙人指", 18)
	fmt.Println(personP)
}
