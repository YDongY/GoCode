package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func SetRequestHeader() string {
	msg := "GET / HTTP/1.1\r\n"
	msg += "Host:www.baidu.com\r\n"
	msg += "Connection:close\r\n"
	msg += "\r\n"

	return msg
}

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		log.Fatalf("Dial error:%v\n", err)
	}

	defer conn.Close()

	headers := SetRequestHeader()

	if _, err = io.WriteString(conn, headers); err != nil {
		log.Fatalf("WriteString error:%v\n", err)
	}

	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("Read error:%v\n", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
