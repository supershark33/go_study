package main

import (
	"encoding/json"
	"fmt"
)

//反射是指在程序运行期对程序本身进行访问和修改的能力。程序在编译时,变量被转换为内存地址，变量名不会被编译器写入到可执行部分。在运行程序时，程序无法获取自身的信息。
//支持反射的语言可以在程序编译期将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息，这样就可以在程序运行期获取类型的反射信息，并且有能力修改它们。
//Go程序在运行期使用reflect包访问程序的反射信息。
//在上- -篇博客中我们介绍了空接口。空接口可以存储任意类型的变量， 那我们如何知道这个空接口保存的数据是什么呢?反射就是在运行时动态的获取一个变量的类型信息和值信息。

//其实就是动态类型（reflect.TypeOf）和动态值(reflect.ValueOf)

//想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，必须传递变量地址才能修改变量值。而反射中使用专有的Elem()| 方法来获取指针对应的值。


import (
	"reflect"
)

func main() {

	str := `{"name":"周林", "age":9000}`
	var p person
	_ = json.Unmarshal([]byte(str), &p)
	fmt.Println(p.Name, p.Age)

	var a float32 = 3.14
	reflectType(a)
	var b int = 9
	reflectType(b)

	p2 := person{
		Name: "刘云克",
		Age:  18,
	}
	reflectType(p2)
	fmt.Println("=================")
	reflectValue2(p2)

	//设置数值
	var c int64 = 9
	//reflectSetValue1(c)
	reflectSetValue2(&c)
	fmt.Println(c)

}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind( )
	switch k {
	case reflect.Int64:
		// v. Int()从反射中获取整型的原始值，然后通过int64( )强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v. Float( )从反射中获取整型的原始值，然后通过float32( )强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v. Float()))
	case reflect.Float64:
		// v. Float( )从反射中获取整型的原始值，然后通过float64( )强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v. Float()))
	}
}

func reflectValue2(x interface{})  {
	valueOf := reflect.ValueOf(x)
	fmt.Println(valueOf)
	fmt.Println(valueOf.Kind())
	typeOf := reflect.TypeOf(x)
	fmt.Println(typeOf)
	fmt.Println(typeOf.Kind())
}

type person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func reflectType(x interface{})  {
	xType := reflect.TypeOf(x)
	fmt.Println(xType)
	fmt.Println(xType.Name())
	fmt.Println(xType.Kind())
}

//通过反射设置变量的值
func reflectSetValue1(x interface{})  {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64{
		v.SetInt(200) //修改的一个副本，reflect会引发panic
	}
}

//通过反射设置变量的值
func reflectSetValue2(x interface{})  {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64{
		v.Elem().SetInt(200) //修改的一个副本，reflect会引发panic
	}
}
