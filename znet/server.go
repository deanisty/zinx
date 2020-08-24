package znet

import (
	"errors"
	"fmt"
	"github.com/deanisty/zinx/ziface"
	"net"
	"time"
)

type Server struct {
	Name string // 服务器的名字
	IPVersion string // 网络版本
	IP string // ip地址
	Port int // 监听端口
}

// =================== 定义当前客户端链接的handler api========================
func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	// 回显业务
	fmt.Println("[Conn Handler] CallBackToClient ... ")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf error : ", err)
		return errors.New("CallBackToClient error")
	}

	return nil
}

func (s *Server) Start() {
	fmt.Printf("Server [%s] starting...\n", s.Name)
	// listening goroutine
	go func() {
		// 1 获取一个tcp的套接字
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve address error : ", err)
			return
		}
		// 2 监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen on address error : ", err)
			return
		}
		fmt.Printf("Listening on %s:%d\n", s.IP, s.Port)

		var cid uint32
		cid = 0
		// 3 处理客户端请求
		for {
			// 3.1 阻塞等待客户端连接
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept connection error : ", err)
				continue
			}
			// 3.2 设置服务器最大连接限制
			// 3.3 处理服务器业务处理
			dealConn := NewConnection(conn, cid, CallBackToClient)
			cid ++
			// 3.4 启动当前链接的处理业务
			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Printf("Server [%s] stopping...", s.Name)
	// 清理资源
}

func (s *Server) Serve() {
	s.Start()
	// post start

	// 阻塞主协程
	for {
		time.Sleep(10 * time.Second)
	}
}

func NewServer (name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      7777,
	}

	return s
}