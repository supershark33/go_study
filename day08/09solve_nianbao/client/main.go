package main

import (
	"fmt"
	"net"
	"zhangyuanbo.com/day08/09solve_nianbao/protocol"
)

func main() {

	conn, err1 := net.Dial("tcp", "127.0.0.1:5000")
	if err1 != nil {
		fmt.Println("dial to server failed, err : ", err1)
		return
	}
	defer conn.Close()
	msg := "hello, hello, how are you?"
	for i:=0;i<20;i++ {
		//调用协议编码
		enMsg, err := protocol.Encode(msg)
		if err != nil {
			fmt.Println("encode err : ", err)
			return
		}
		conn.Write(enMsg)
		//time.Sleep(time.Second)
	}
}
