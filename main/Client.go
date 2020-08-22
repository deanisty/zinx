package main

import(
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("Client start...")

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("Start error : ", err)
		return
	}

	for {
		_, err = conn.Write([]byte("hello Zinx"))
		if err != nil {
			fmt.Println("Write error : ", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read error : ", err)
			return
		}

		fmt.Printf("Server echo : %s, cnt = %d\n", buf, cnt)

		time.Sleep(1 * time.Second)
	}
}
