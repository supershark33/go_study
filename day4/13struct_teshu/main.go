package main

import (
	"fmt"
)

func main() {

	//1、结构体匿名字段,可以看懂，但是不要用
	p := person{
		18, "liuyunke",
	}
	fmt.Println(p)
	fmt.Println(p.int)
	fmt.Println(p.string)

	//2、结构体嵌套
	//可以认为，类型就是名称
	//相同的类型只能出现一个
	s := student{
		age:  15,
		name: "刘云克",
		addr: address{
			province: "河南",
			city:     "三门峡",
		},
	}
	fmt.Println(s)

	//2、结构体匿名嵌套
	s2 := student2{
		int:  19,
		name: "仙人指",
		address: address{
			province: "山西",
			city:     "大同",
		},
	}
	fmt.Println(s2)
	fmt.Println(s2.int)
	fmt.Println(s2.name)
	fmt.Println(s2.province) //先找自己结构体有没有这个字段，如果没有就到匿名结构体里面去找
	fmt.Println(s2.city)
	fmt.Println(s2.address.province)
	fmt.Println(s2.address.city)
}

//匿名字段
type person struct {
	int
	string
}

//正常嵌套
type student struct {
	age int
	name string
	addr address
}

type address struct {
	province string
	city string
}

//匿名嵌套
type student2 struct {
	int
	name string
	address
}




