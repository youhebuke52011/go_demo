package main

import "net"

func handleConn(conn net.Conn) {
	defer conn.Close()
	//for {
	//	//conn.Read()
	//}
}

func main() {
	listener, err := net.Listen("tcp", "9999")
	if err != nil{
		return
	}

	for  {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		go handleConn(conn)
	}
}
