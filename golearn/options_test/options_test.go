package optionstest

import (
	"fmt"
	"testing"
)

func TestFuncOptions(t *testing.T) {
	s := NewServer("192.168.100.213", 65, Withtimeout(1000), WithName("lilei"))

	fmt.Printf("s.name: %v\n", s.name)
}

type Server struct {
	ip      string
	port    int
	timeout int
	name    string
}

type Options func(s *Server)

func Withtimeout(timeout int) Options {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithName(name string) Options {
	return func(s *Server) {
		s.name = name
	}
}

func NewServer(ip string, port int, options ...Options) *Server {
	s := &Server{
		ip:      ip,
		port:    port,
		timeout: 20,
		name:    "defalt",
	}

	for _, o := range options {
		o(s)
	}

	return s
}
