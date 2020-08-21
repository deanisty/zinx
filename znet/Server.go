package znet

import (
	"fmt"
	"net"
)

type Server struct {
	Name string
	IPVersion string
	IP string
	Port int
}

func (s *Server) Start() {
	fmt.Printf("Server [%s] starting, listening on IP: %s, Port: %d\n", s.Name, s.IP, s.Port)
	net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
}

func (s *Server) Stop() {

}

func (s *Server) Server() {

}