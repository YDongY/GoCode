package main

import (
	"fmt"
	"github.com/07_网络编程/TCPProject/server/model"
	"net"
	"time"
)

func process(conn net.Conn) {

	defer conn.Close()
	p := &Processor{Conn: conn}
	err := p.process2()
	if err != nil {
		return
	}

}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	// 服务器启动，初始化连接池
	initPool(":6379", 16, 0, 300*time.Second)
	initUserDao()

	fmt.Println("服务器在【9999】端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listen.Close()
	for {
		fmt.Println("等待客户端连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err:", err)
		}
		go process(conn)
	}

}
