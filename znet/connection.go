package znet

import (
	"errors"
	"fmt"
	"github.com/deanisty/zinx/ziface"
	"io"
	"net"
)

type Connection struct {
	// 当前连接的tcp套接字
	Conn *net.TCPConn
	// 当前连接的id 全局唯一
	ConnID uint32
	// 连接是否关闭
	isClosed bool
	// 该连接的处理方法router
	Router ziface.IRouter
	// 通知当前连接关闭的channel
	ExitBuffChan chan bool
}

func NewConnection(conn *net.TCPConn, connId uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn: conn,
		ConnID: connId,
		isClosed: false,
		Router: router,
		ExitBuffChan: make(chan bool, 1),
	}

	return c
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running")
	defer fmt.Println(c.RemoteAddr().String(), " conn reader exit!")
	defer c.Stop()

	for {
		// 封包拆包对象
		dp := NewDataPack()
		// 读取header
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head error : ", err)
			c.ExitBuffChan <- true
			continue
		}
		// 拆包
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error : ", err)
			c.ExitBuffChan <- true
			continue
		}
		// 根据 dataLen 读取data
		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read data error : ", err)
				c.ExitBuffChan <- true
				continue
			}
		}
		msg.SetData(data)
		// 构造客户端请求 request
		req := Request{
			conn: c,
			data: msg,
		}
		// 从 router 中找到注册绑定 Conn 对应的 Handler
		go func (request ziface.IRequest) {
			// 执行注册路由的方法
			c.Router.PreHandler(request)
			c.Router.Handler(request)
			c.Router.PostHandler(request)
		}(&req)
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

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

/**
 发送数据到客户端
 */
func (c *Connection) SendMessage(msgId uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("connection closed when send msg")
	}
	// 封包
	dp := NewDataPack()
	msg, err := dp.Pack(NewMessagePacket(msgId, data))
	if err != nil {
		fmt.Println("Pack msg error : id ", msgId)
		return errors.New("pack message error")
	}
	// 写回客户端
	if _, err = c.Conn.Write(msg); err != nil {
		fmt.Println("Write message error : id ", msgId)
		c.ExitBuffChan <- true
		return errors.New("conn write error")
	}

	return nil
}