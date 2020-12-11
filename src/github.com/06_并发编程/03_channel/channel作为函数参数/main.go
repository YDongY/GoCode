package main

import (
	"fmt"
)

func channelReader(messages <-chan string) {
	msg := <-messages
	fmt.Println(msg)
}

func channelWriter(message chan<- string) {
	message <- "hello world"
}

func channelReaderAndWriter(message chan string) {
	msg := <-message
	fmt.Println(msg)
	message <- "hello world"
}

func main() {
	// <-chan 只读
	// chan<- 只写
	// chan   读写
}
