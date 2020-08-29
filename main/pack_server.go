package main

import (
	"fmt"
	"github.com/deanisty/zinx/znet"
	"io"
	"net"
)

func main() {
	// 创建 TCP 监听
	listenner, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("server listen error : ", err)
		return
	}

	// 接收客户端请求
	for {
		conn, err := listenner.Accept()
		if err != nil {
			fmt.Println("server accept error : ", err)
			continue
		}

		// 处理请求
		go func(conn net.Conn) {
			// 创建包
			dp := znet.NewDataPack()
			for {
				// 读取head
				headData := make([]byte, dp.GetHeadLen())
				_, err := io.ReadFull(conn, headData)
				if err != nil {
					fmt.Println("read head error : ", err)
					break
				}
				// head 拆包
				msgHead, err := dp.Unpack(headData)
				if err != nil {
					fmt.Println("server unpack error : ", err)
					break
				}

				msg := msgHead.(*znet.Message)

				if msgHead.GetDataLen() > 0 {
					// data有数据
					msg.Data = make([]byte, msg.GetDataLen())

					_, err := io.ReadFull(conn, msg.Data)
					if err != nil {
						fmt.Println("server read data error : ", err)
						return
					}
				}

				fmt.Println("===> Recv Msg: ID=", msg.Id, "len=", msg.GetDataLen(), "data=", string(msg.Data))
			}
		}(conn)
	}
}
