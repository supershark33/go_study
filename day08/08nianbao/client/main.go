package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	conn, err1 := net.Dial("tcp", "127.0.0.1:40000")
	if err1 != nil {
		fmt.Println("dial to server failed, err : ", err1)
		return
	}
	defer conn.Close()
	msg := "hello, hello, how are you?"
	for i:=0;i<20;i++ {
		conn.Write([]byte(msg))
		time.Sleep(time.Second)
	}
}
