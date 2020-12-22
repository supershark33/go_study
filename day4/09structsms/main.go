package main

import (
	"fmt"
	"os"
)

var (
	students map[int64]*student
)

func main() {

	if(students == nil){
		students = make(map[int64]*student, 64)
	}

	for{
		//打印菜单
		fmt.Println("欢迎光顾学生管理系统！")
		fmt.Println(`
		1、查询所有学生
		2、新增学生
		3、删除学生
		4、退出
	`)
		fmt.Println("请输入你要干啥...")
		//等待用户选择
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d的选项\n", choice)
		//执行对应函数
		switch choice {
		case 1:
			showAllStudects()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("未知类型，请重新输入~")
		}
	}

}

func deleteStudent() {
	fmt.Println("请输入要删除的学生学号...")
	var id int64
	fmt.Scanln(&id)
	delete(students, id)
	fmt.Printf("学号%d的学生删除成功", id)
}

func addStudent() {
	fmt.Println("请输入学生学号...")
	var id int64
	fmt.Scanln(&id)
	fmt.Println("请输入学生姓名...")
	var name string
	fmt.Scanln(&name)
	newStu:= newStudent(id, name)
	students[id] = newStu
}

func newStudent(id int64, name string) *student {
	return &student{
		id: id,
		name: name,
	}
}

func showAllStudects() {
	if(students == nil){
		return
	}
	for key,value := range students{
		fmt.Println(key,value)
	}
}

type student struct {
	id int64
	name string
}


