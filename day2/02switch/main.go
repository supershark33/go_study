package main

import "fmt"

//switch 简化大量判断，一个变量和给定值作比较
func main() {

	//标准方式
	for n:=0;n<10;n++{
		switch n {
		case 1:
			fmt.Println("大拇指")
		case 2:
			fmt.Println("食指")
		case 3:
			fmt.Println("中指")
		case 4:
			fmt.Println("无名指")
		case 5:
			fmt.Println("小拇指")
		default:
			fmt.Println("未知的数字")
		}
	}

	//变种二
	age := 35
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age >= 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}

	//变种三


	//变种四

}

//func switchFun(int n)  {
//
//}
