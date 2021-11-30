package main

import (
	"bufio"
	"fmt"
	"golab/packages-example/net/tcp_sticky_package/proto"
	"io"
	"log"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf [1024]byte
	for {
		n, err := reader.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Read error:%v\n", err)
			break
		}
		fmt.Println(string(buf[:n]))
	}
}

func processProto(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Printf("Read error:%v\n", err)
			return
		}
		fmt.Println(msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("Listen error:%v\n", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error:%v\n", err)
			continue
		}
		// go process(conn)
		go processProto(conn)
	}
}
