package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	udp, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8080,
	})
	if err != nil {
		log.Fatalf("DialUDP error:%v\n", err)
	}
	defer udp.Close()

	// 发送数据
	_, err = udp.Write([]byte("Hello Server"))
	if err != nil {
		log.Fatalf("Write error:%v\n", err)
	}

	// 接收数据
	buf := make([]byte, 1024)
	n, addr, err := udp.ReadFromUDP(buf)
	if err != nil {
		log.Fatalf("ReadFromUDP error:%v\n", err)
	}
	fmt.Printf("[%s] >>> %s\n", addr, string(buf[:n]))
}
