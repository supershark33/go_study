package main

import (
	"encoding/json"
	"fmt"
)

//结构体与json
//1、把go语言中的结构体 -> 转换为json格式的字符串
//2、把json格式的字符串 -> 转换为go语言的结构体
//3、标识符的可见性以及字段的可见性，首字符大小写的问题

type person struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age int `json:"age"`
}

func main() {

	p := person{
		Name: "刘云克",
		Age:  250,
	}

	//序列化
	marshal, err := json.Marshal(p)
	if(err != nil){
		fmt.Printf("转换错误：%v", err)
		return
	}
	fmt.Println(string(marshal))

	//反序列化
	str := `{"name":"理想","age":18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2)
	fmt.Println(p2)
	fmt.Printf("%v\n", p2)
	fmt.Println(p2.Name, p2.Age)
}
