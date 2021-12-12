package main

import (
	"golab/packages-example/net/tcp_sticky_package/proto"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalf("Dial error:%v\n", err)
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello Golang`

		data, err := proto.Encode(msg)
		if err != nil {
			log.Fatalf("Encode error:%v\n", err)
			return
		}
		conn.Write(data)

		// conn.Write([]byte(msg))
	}
}
