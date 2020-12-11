package main

import "fmt"

func slowFunc(c chan string) {
	for msg := range c {
		fmt.Println(msg)
	}
}

func main() {
	c := make(chan string, 2) // 缓冲l通道最多只能存储指定数量的消息

	c <- "hello"
	c <- "world"

	close(c) // 关闭 channel，之后只能读不能写

	fmt.Println("---main---")
	slowFunc(c) // channel 是引用类型

	// c <- "close after write" // panic: send on closed channel
}
