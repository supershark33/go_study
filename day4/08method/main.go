package main

import "fmt"

//方法是作用域特定类型的函数
//函数是任意都可以调用的
func main() {

	d := newDog("刘云克", 20)
	d.wang()
	fmt.Println(d)
	d.wangP()
	fmt.Println(d)
}

//go语言首字母大写 标识对外部可见 是共有的
//如果首字母是小写的，则在外面的包是找不到的，对外是不见见的
type Dog struct {
	name string
	age int
}

func newDog(name string, age int) Dog {
	return Dog{
		name: name,
		age: age,
	}
}


//1.需要修改接收者中的值
//2.接收者是拷贝代价比较大的大对象
//3.保证一致性,如果有某个方法使用了指针接收者,那么其他的方法也应该使用指针接收者。
//值接受者
func (d Dog)wang() {
	fmt.Printf("%s汪汪汪~", d.name)
	d.age++
}

//指针接受者
func (d *Dog) wangP() {
	d.age++
}

