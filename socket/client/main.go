package main

import (
	"log"
	"net"
	"time"
)

func client1() {
	conn, err := net.DialTimeout("tcp", ":9999", 2*time.Second)
	if err != nil {
		log.Printf("err:%v", err)
		return
	}
	defer conn.Close()
	log.Printf("client send!")
}

func client2(i int) net.Conn {
	//conn, err := net.DialTimeout("tcp", ":9999", 2*time.Second)
	conn, err := net.Dial("tcp", ":9999")
	if err != nil {
		log.Printf("err:%v", err)
		return nil
	}
	//defer conn.Close()
	//log.Printf("client send!")
	log.Println(i, ":connect to server ok")
	return conn
}

func client3(data string) {
	conn, err := net.DialTimeout("tcp", ":9999", 2*time.Second)
	if err != nil {
		log.Printf("err:%v", err)
		return
	}
	defer conn.Close()
	n, err := conn.Write([]byte(data))
	if err != nil {
		log.Printf("client write data error:%v",err)
		return
	}
	log.Printf("client write %d bytes, content is %s\n", n, data)
}


func client4() {
	conn, err := net.DialTimeout("tcp", ":9999", 2*time.Second)
	if err != nil {
		log.Printf("err:%v", err)
		return
	}
	defer conn.Close()
	data := make([]byte, 65536)
	n, err := conn.Write(data)
	if err != nil {
		log.Printf("client write data error:%v",err)
		return
	}
	log.Printf("client write %d bytes, content is %s\n", n, data)
}

func main() {
	//client1()

	//for i := 0; i < 1000; i++ {
	//	conn := client2(i)
	//	if conn != nil {
	//
	//	}
	//}
	//time.Sleep(1 * time.Hour)

	//client3(os.Args[1])

	client4()
}
