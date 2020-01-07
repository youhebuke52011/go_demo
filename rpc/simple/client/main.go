package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:6969")
	if err != nil {
		log.Fatal("dialing:", err)
		return
	}

	var reply string
	err = client.Call("HelloService.Hello", "rpc test", &reply)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("rpc client call service response:%s\n", reply)
}
