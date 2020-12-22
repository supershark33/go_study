package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"zhangyuanbo.com/day08/09solve_nianbao/protocol"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:5000")
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
		go process2(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var readBuf [1024]byte //直接读取1024会导致粘包问题
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

func process2(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	//这里不进行固定长度读取，而是根据长度协议来
	for {
		recvStr, err := protocol.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode err : ", err)
			return
		}
		fmt.Println("收到client发来的数据: ", recvStr)
	}
}
