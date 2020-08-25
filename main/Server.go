package main

import (
	"fmt"
	"github.com/deanisty/zinx/ziface"
	"github.com/deanisty/zinx/znet"
)

// 自定义路由
type PingRouter struct {
	znet.BaseRouter //继承基础路由
}

func (this *PingRouter) PreHandler(request ziface.IRequest) {
	fmt.Println("Call PingRouter PreHandler")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping ...\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

func (this *PingRouter) Handler(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handler")
	_,err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

func (this *PingRouter) PostHandler (request ziface.IRequest)  {
	fmt.Println("Call PingRouter PostHandler")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping ...\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

func main() {
	s := znet.NewServer()
	s.AddRouter(&PingRouter{})
	s.Serve()
}
