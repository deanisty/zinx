package main

import (
	"fmt"
	"github.com/deanisty/zinx/znet"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dial error : ", err)
		return
	}

	dp := znet.NewDataPack()
	msg1 := &znet.Message{
		Id:      0,
		Data:    []byte{'h', 'e', 'l', 'l', 'o'},
		DataLen: 5,
	}
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client pack msg1 error: ", err)
		return
	}

	msg2 := &znet.Message{
		Id:      1,
		Data:    []byte{'w', 'o', 'r', 'l', 'd', '!', '!'},
		DataLen: 7,
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client pack msg2 error: ", err)
	}

	sendData1 = append(sendData1, sendData2...)
	conn.Write(sendData1)

	select { }
}
