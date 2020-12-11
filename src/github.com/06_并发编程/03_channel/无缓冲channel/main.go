package main

import (
	"fmt"
	"time"
)

// 创建一个无缓冲通道
var c = make(chan string)

func slowFunc() {
	time.Sleep(time.Second * 2)
	c <- "slowFunc()"
}

func main() {
	go slowFunc()
	msg := <-c // 管道没有数据，阻塞等待
	fmt.Println(msg)

}
