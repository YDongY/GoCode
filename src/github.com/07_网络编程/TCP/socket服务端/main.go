package main

import (
	"fmt"
	"io"
	"net"
)

func handler(conn net.Conn) {
	// 循环接收数据
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Printf("[%s] client close \n", conn.RemoteAddr().String())
			return
		}
		fmt.Println(string(buf[:n]))
	}
}

func main() {
	fmt.Println("开始监听......")
	listener, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		fmt.Println("Listen error")
	}

	defer listener.Close()

	// 等待连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept error")
		}
		fmt.Printf("from %s conn success \n", conn.RemoteAddr().String())
		go handler(conn)
	}

}
