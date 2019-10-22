package main

import (
	"log"
	"net"
	"time"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		// read from the connection
		buf := make([]byte, 10)
		log.Println("start to read from conn")
		conn.SetReadDeadline(time.Now().Add(time.Microsecond * 10))
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("server read data err:%v", err)
			return
		}
		log.Printf("server read %d bytes, content is %s\n", n, string(buf[:n]))

		// write to the connection
		//data := "你好"
		//n, err = conn.Write([]byte(data))
		//if err != nil {
		//	log.Printf("server write data error:%v",err)
		//	return
		//}
		//log.Printf("server write %d bytes, content is %s\n", n, data)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		return
	}
	defer listener.Close()

	i := 0
	for {
		time.Sleep(10 * time.Second)
		conn, err := listener.Accept()
		if err != nil {
			break
		}
		i++
		log.Printf("%d: accept a new connection\n", i)
		go handleConn(conn)
	}
}
