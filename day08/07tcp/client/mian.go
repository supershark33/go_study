package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)


func main() {
	//与客户端简历连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err!= nil {
		fmt.Println("connect to server failed, err:", err)
		return
	}
	//发送数据
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("请说话...")
		msg, err := reader.ReadString('\n')
		if err != nil || msg == "exit" {
			break
		}
		_, err1 := conn.Write([]byte(msg))
		if err1 != nil {
			fmt.Println("write failed, err:", err1)
		}
	}
	conn.Close()
}
