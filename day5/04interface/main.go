package main

import "fmt"

func drive(c car)  {
	c.run()
}

type car interface {
	run()
}

type byd struct {
	name string
}

func (b byd)run()  {
	fmt.Printf("%s速度59迈~\n", b.name)
}


type bmw struct {
	name string
}

func (b bmw)run()  {
	fmt.Printf("%s速度120迈\n", b.name)
}

type haval struct {
	name string
}

func (h haval) run() {
	fmt.Printf("%s速度10迈\n", h.name)
}

func main() {
	bydCar:= byd{
		name: "比亚迪",
	}

	bmwCar := bmw{
		name: "奔驰",
	}

	havalCar := haval{
		name: "哈弗",
	}

	drive(bydCar)
	drive(bmwCar)
	drive(havalCar)
}
