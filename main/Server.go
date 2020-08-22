package main

import (
	"github.com/deanisty/zinx/znet"
)

func main() {
	s := znet.NewServer("zinx v1.0")
	s.Serve()
}
