package main

import (
	"fmt"
	"reflect"
)

//反射是-一个强大并富有表现力的工具，能让我们写出更灵活的代码。但是反射不应该被滥用,原因有以下三个。
//1.基于反射的代码是极其脆弱的，反射中的类型错误会在真正运行的时候才会引发panic,那很可能是在代码写完的很长时间之后。
//2.大量使用反射的代码通常难以理解。
//3.反射的性能低下，基于反射实现的代码通常比正常代码运行速度慢一到两个数量级。

func main() {

	stu1 := student{
		Name:  "花小猪",
		Score: 100,
	}
	stuType := reflect.TypeOf(stu1)
	fmt.Println(stuType)
	fmt.Println(stuType.Name(), stuType.Kind()) //student struct

	//通过for循环获取结构体所有的fields
	for i:=0;i< stuType.NumField();i++{
		field := stuType.Field(i)
		fmt.Printf("name:%s, index:%d, type:%v, json.tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("zhoulin"))
	}

	//指定名称获取结构体信息
	if scroeField, ok := stuType.FieldByName("Score"); ok{
		fmt.Printf("name:%s, index:%d, type:%v, json.tag:%v\n", scroeField.Name, scroeField.Index, scroeField.Type, scroeField.Tag.Get("zhoulin"))
	}

	stuValue := reflect.ValueOf(stu1)
	fmt.Println(stuValue)
	name := stuValue.FieldByName("Name")
	score := stuValue.FieldByName("Score")
	fmt.Println(name, name.Type())
	fmt.Println(score, score.Type())
}

type student struct {
	Name string `json:"name" zhoulin:"嘿嘿嘿"`
	Score int `json:"score" zhoulin:"哈哈哈"`
}


