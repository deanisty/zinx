package znet

import "net"

type Connection struct {
	// 当前连接的tcp套接字
	Conn *net.TCPConn
	// 当前连接的id 全局唯一
	ConnID uint32
	// 连接是否关闭
	isClosed bool
	// handler
	handleAPI ziface.HandFunc
	// 通知当前连接关闭的channel
	ExitBuffChan chan bool
}

func NewConnection(conn *net.TCPConn, connId uint32, callback_api ziface.HandFunc) *Connection {
	c := &Connection{
		Conn: conn,
		ConnID: connId,
		isClosed: false,
		handleAPI: callback_api,
		ExitBuffChan: make(chan bool, 1),
	}

	return c
}