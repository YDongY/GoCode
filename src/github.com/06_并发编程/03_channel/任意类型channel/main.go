package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	c := make(chan interface{}, 5)

	p := Person{name: "jack", age: 20}

	c <- "hello world"
	c <- 100
	c <- p
	c <- false
	close(c) // 遍历channel之前，需要关闭，否则 fatal error: all goroutines are asleep - deadlock!
	for {
		msg := <-c
		if msg == nil {
			break
		}
		fmt.Println(msg)
	}
}
