package main

import (
	"fmt"
	"time"
)

func ping1(c chan string) {
	time.Sleep(time.Second * 1)
	c <- "ping1"
}

func ping2(c chan string) {
	time.Sleep(time.Second * 2)
	c <- "ping2"
}

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go ping1(channel1)
	go ping2(channel2)

	// 收到一条消息，解阻塞，执行一条 case 其他消息将被丢弃
	select {
	case msg1 := <-channel1:
		fmt.Println("recv msg", msg1)
	case msg2 := <-channel2:
		fmt.Println("recv msg2:", msg2)
	case <-time.After(500 * time.Millisecond): // 指定在 O.Ss 内没有收到消息时将采取的措施
		fmt.Println("no messages received . giving up . ")
	}
}
