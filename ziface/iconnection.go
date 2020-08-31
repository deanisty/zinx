package ziface

import "net"

// 定义连接接口
type IConnection interface {
	// 启动连接
	Start()
	// 停止连接
	Stop()
	// 获取socket
	GetTCPConnection() *net.TCPConn
	// 获取当前连接ID
	GetConnID() uint32
	// 获取远程客户端地址
	RemoteAddr() net.Addr
	// 发送 message 数据给客户端
	SendMessage(msgId uint32, data []byte) error
}
