package main

import (
	"fmt"
	"net"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//1、本地端口启动服务
	listener, err1 := net.Listen("tcp", "127.0.0.1:20000")
	if err1 != nil{
		fmt.Println("start tcp server on 127.0.0.1:8888 failed, err:", err1)
		return
	}
	wg.Add(1)
	fmt.Println("服务端口20000启动成功，等待连接...")
	//2、等待别人来建立连接
	for {
		conn, err2 := listener.Accept()
		if err2 != nil {
			fmt.Println("accept failed, err:", err2)
			return
		}
		//3、与客户端通信
		go process(conn)
	}
	wg.Wait()
}

func process(conn net.Conn) {
	var temp [128]byte
	for {
		n, err3 := conn.Read(temp[:])
		if err3 != nil {
			fmt.Println("read from conn failed, err:", err3)
			break
		}
		fmt.Println(string(temp[:n]))
	}
	conn.Close()
}