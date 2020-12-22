package main

import "fmt"

//结构体内部是没有继承的，但是可以模拟类似的功能
func main() {

	d := dog{
		animal: animal{
			name: "刘云克",
		},
		feet: 4,
	}

	d.wang()
	d.move()
}

type animal struct {
	name string
}

func (a animal) move() {
	fmt.Printf("%s会动\n", a.name)
}

type dog struct {
	feet uint8
	animal
}

func (d dog) wang() {
	fmt.Printf("%s汪汪叫~\n", d.name)
}




