package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":6969")
	if err != nil{
		log.Fatal("ListenTCP error:", err)
		return
	}

	conn, err := listener.Accept()
	if err != nil{
		log.Fatal("Accept error:", err)
		return
	}

	rpc.ServeConn(conn)
}
