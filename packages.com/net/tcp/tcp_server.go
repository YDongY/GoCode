package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Listen error:%v\n", err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error:%v\n", err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 128)
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("Accept error:%v\n", err)
			break
		}
		fmt.Println(string(buf[:n]))
	}
}
