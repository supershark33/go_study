package main

import "fmt"

//空接口一般不用起名称，因为所有的结构体都实现了空接口，因此空接口可以用来接所有结构体的引用

//空接口作为map的参数

//空接口作为函数参数
func show(a interface{})  {
	fmt.Printf("type:%T value:%v\n", a, a)
}

//类型断言 既然空接口可以存放任意类型的值，那我们如何获取其存储的具体数据呢？类型断言

func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "周林"
	m1["age"] = 9000
	m1["merried"] = true
	m1["hobby"] = [...]string{"唱", "跳", "rap"}
	fmt.Println(m1)

	show(false)
	show(nil)
	show(m1)

	//类型断言
}
