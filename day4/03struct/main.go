package main

import "fmt"

//结构体
//其中:
//类型名:标识自定义结构体的名称,在同一个包内不能重复。
//字段名:表示结构体字段名。结构体中的字段名必须唯一
//字段类型:表示结构体字段的具体类型
type person struct{
	name string
	age int
	gender string
	hobby []string
}

//p是一个副本
func f1(p person)  {
	p.name = "仙人指"
}

//p是一个引用
func f2(pp *person)  {
	(*pp).name = "刘云克"
	//如果pp是指针类型的，在go内部是有语法唐的，可以不用（*pp），直接pp就好
	pp.name = "刘云克"
}

func main() {

	var p person
	p.name = "刘云克"
	p.age = 21
	p.gender = "男"
	p.hobby = []string{"篮球","足球", "双色球"}
	fmt.Println(p)
	//访问字段
	fmt.Printf("%T\n", p)
	fmt.Println(p.name)

	//匿名结构体 多用于临时场景
	var s struct{
		age int
		name string
	}
	s.age = 11
	s.name = "嘿嘿嘿"
	fmt.Printf("type:%T value:%v\n", s)

	//结构体是值类型的 go语言里面 函数传递的永远都是拷贝
	var s1 person
	s1.name = "仙人指"
	fmt.Println(s1)
	f1(s1)
	fmt.Println(s1)
	//值类型数据传递内存地址
	f2(&s1)
	fmt.Println(s1)

	//我们还可以用new来对结构体进行实例化，new是对基本类型进行实例化的，返回的是指针类型，但是要考虑语法唐哦，
	//用起来还真差不多，而make只能针对slice map 和 管道，返回的是值
	//结构体指针1
	p2 := new(person)
	p2.name = "张拾玥"
	fmt.Printf("%T\n", p2) //*main.person
	fmt.Printf("%v\n", p2)
	fmt.Printf("%p\n", &p2)
	//fmt.Printf("%T\n", s1)
	//fmt.Printf("%v\n", s1)

	//结构体指针2
	//2.1 key-value初始化，不一定全传
	var p3 = &person{
		name : "张三",
		gender: "男",
	}
	fmt.Printf("%T\n", p3) //main.person
	fmt.Printf("%v\n", p3)
	fmt.Printf("%p\n", &p3)
	fmt.Printf("%p\n", &p3)

	//2.2使用值列表进行初始化，指的顺序要和字段的定义顺序一致
	p4 := &person{
		"克克",108,"中性", []string{"看黄色小说"},
	}
	fmt.Printf("%T\n", p4) //*main.person
	fmt.Printf("%v\n", p4)
	fmt.Printf("%p\n", &p4)

	//结构体占用一款连续的内存
	type x struct {
		a int8
		b int8
		c int8
	}
	m := x{
		int8(10),int8(20),int8(30),
	}
	fmt.Printf("%p\n", &(m.a))
	fmt.Printf("%p\n", &(m.b))
	fmt.Printf("%p\n", &(m.c))



}