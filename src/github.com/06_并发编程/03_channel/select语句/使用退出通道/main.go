package main

import (
	"fmt"
	"time"
)

func sender(c chan string) {

	t := time.NewTicker(1 * time.Second)
	for {
		c <- "---sender--"
		<-t.C
	}
}

func main() {
	// 些情况下， 不确定 select 语句该在何时返回，因此不能使用定时器

	message := make(chan string)
	stop := make(chan bool)

	go sender(message)

	go func() {
		time.Sleep(time.Second * 2)
		stop <- true
	}()

	for {
		select {
		case <-stop:
			return
		case msg := <-message:
			fmt.Println(msg)
		}
	}
}
