package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	udp, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8080,
	})

	if err != nil {
		log.Fatalf("ListenUDP error:%v\n", err)
	}

	for {
		buf := make([]byte, 1024)
		n, addr, err := udp.ReadFromUDP(buf)
		if err != nil {
			log.Printf("ReadFromUDP error:%v\n", err)
			continue
		}
		fmt.Printf("[%s] >>> %s\n", addr, string(buf[:n]))

		// 发送数据
		_, err = udp.WriteToUDP([]byte("Hello Client"), addr)
		if err != nil {
			log.Printf("WriteToUDP error:%v\n", err)
			continue
		}
	}

}
