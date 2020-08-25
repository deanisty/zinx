package znet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func ClientTest() {
	fmt.Println("Client test ... start")
	// 等待3秒 让服务端启动
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start error : ", err)
		return
	}

	for {
		_, err = conn.Write([]byte("hello Zinx"))
		if err != nil {
			fmt.Println("client write error : ", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error : ", err)
			return
		}

		fmt.Printf("Server echo back : %s, cnt = %d\n", buf, cnt)

		time.Sleep(1 * time.Second)
	}
}

func TestServer(t *testing.T) {
	s := NewServer()

	go ClientTest()

	s.Serve()
}
