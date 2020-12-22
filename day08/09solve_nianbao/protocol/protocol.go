package protocol

import (
	"bufio"
	"bytes"
	"encoding/binary"
)
//出现”粘包”的关键在于接收方不确定将要传输的数据包的大小，因此我们可以对数据包进行封包和拆包的操作。
//封包:封包就是给一段数据加 上包头,这样一来数据包就分为包头和包体两部分内容了(过滤非法包时封包会加入”包尾”内容)。包头部分的长度是固定的，并且它存储了包体的长度，根据包头长度固定以及包头中含有包体长度的变量就能正确的拆分出一个完整的数据包。
//我们可以自己定义-一个协议，比如数据包的前4个字节为包头，里面存储的是发送的数据的长度。

//发送消息的时候将信息编码
func Encode(message string) ([]byte, error) {
	//读取消息长度，转换成int32类型（4个字节固定长度）
	var length = int32(len(message))
	pkg := new(bytes.Buffer)
	//写入消息头
	err1 := binary.Write(pkg, binary.LittleEndian, length)
	if err1 != nil {
		return nil, err1
	}
	//写入消息体
	err2 := binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err2 != nil{
		return nil, err2
	}
	return pkg.Bytes(), nil
}

func Decode(reader *bufio.Reader) (string, error){
	//读取消息长度
	lengthBytes, _ := reader.Peek(4) //读取前4个字节的数据
	buffer := bytes.NewBuffer(lengthBytes)
	var length int32
	err1 := binary.Read(buffer, binary.LittleEndian, &length)
	if err1 != nil {
		return "", err1
	}
	//buffered返回缓冲中咸鱼的可以读取的字节数
	if int32(reader.Buffered()) < length+4 {
		return  "", err1
	}
	//读取整整的消息数据
	pack := make([]byte, int(4+length))
	_, err2 := reader.Read(pack)
	if err2 != nil {
		return "", err2
	}
	return string(pack[4:]), nil
}