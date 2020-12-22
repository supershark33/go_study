package main

import "fmt"

//类型断言
func assign(a interface{})  {
	fmt.Printf("%T\n", a)
	s,ok := a.(string)
	if !ok {
		fmt.Println("类型转换失败")
		return
	}
	fmt.Println(s)
}

func assign2(a interface{})  {
	fmt.Printf("%T\n", a)
	switch t := a.(type){
	case string:
		fmt.Printf("我是一个字符串：", t)
	case int:
		fmt.Printf("我是一个int：", t)
	case float32:
		fmt.Printf("我是一个float32：", t)
	default:
		fmt.Println("未知类型")
	}

}

func main() {

	assign(100)

	assign2("xxx")
}


