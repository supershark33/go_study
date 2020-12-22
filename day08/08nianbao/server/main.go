package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:40000")
	if err != nil {
		fmt.Println("listen failed, err : ", err)
		return
	}
	defer listen.Close()
	//循环进行accpet，main的gorouting不会退出
	for {
		conn, err1 := listen.Accept()
		if err1 != nil {
			fmt.Println("accept failed, err : ", err1)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var readBuf [1024]byte
	for {
		n, err := reader.Read(readBuf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read failed, err : ", err)
			break
		}
		recvStr := string(readBuf[:n])
		fmt.Println("收到client发来的数据: ", recvStr)
	}
	//read, err2 := conn.Read(readBuf[:])
	//if err2 != nil {
	//	fmt.Println("read failed, err : ", err2)
	//}
	//fmt.Println(string(readBuf[:read]))
}
