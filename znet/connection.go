package znet

import (
	"fmt"
	"github.com/deanisty/zinx/ziface"
	"net"
)

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

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running")
	defer fmt.Println(c.RemoteAddr().String(), " conn reader exit!")
	defer c.Stop()

	for {
		// 读取数据到buf中 阻塞
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf error : ", err)
			c.ExitBuffChan <- true
			continue
		}
		// 调用当前链接业务方法
		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("connId : ", c.ConnID, " handler is error")
			c.ExitBuffChan <- true
			return
		}
	}
}

func (c *Connection) Start() {
	// 开启
	go c.StartReader()

	//for {
	//	select {
	//	case <- c.ExitBuffChan:
	//		// 得到推出消息，不再阻塞
	//		return
	//	}
	//}
}

func (c *Connection) Stop() {
	// 当前连接已经关闭
	if c.isClosed == true {
		return
	}

	// TODO Connection.Stop() 如果用户注册了该链接的关闭回调业务，那么在此刻应该显示调用

	// 关闭socket链接
	c.isClosed = true

	// 通知从缓冲队列读数据的业务，该链接已经关闭
	c.ExitBuffChan <- true

	// 关闭channel
	close(c.ExitBuffChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnId() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}