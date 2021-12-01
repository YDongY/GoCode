package main

import (
	"fmt"
	"time"
)

func TestSelect1() {
	// select 可以同时监听一个或多个 channel，直到其中一个 channel ready
	output1 := make(chan string)
	output2 := make(chan string)

	go test1(output1)
	go test2(output2)

	select {
	case s1 := <-output1:
		fmt.Println("s1=", s1)
	case s2 := <-output2:
		fmt.Println("s2=", s2)
	}
}

func test1(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test1"
}
func test2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "test2"
}

func TestSelect2() {
	// 如果多个 channel 同时 ready，则随机选择一个执行
	// 创建2个管道
	ch1 := make(chan int, 1)
	ch2 := make(chan string, 1)
	go func() {
		ch1 <- 1
	}()
	go func() {
		ch2 <- "hello"
	}()
	select {
	case value := <-ch1:
		fmt.Println("int:", value)
	case value := <-ch2:
		fmt.Println("string:", value)
	}
	fmt.Println("main 结束")
}

func TestSelect3() {
	// 用于判断管道是否存满
	// 创建管道
	output1 := make(chan string, 10)
	// 子协程写数据
	go func() {
		for {
			select {
			// 写数据
			case output1 <- "hello":
				fmt.Println("write hello")
			default:
				fmt.Println("channel full")
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()
	// 取数据
	for s := range output1 {
		fmt.Println("res:", s)
		time.Sleep(time.Second)
	}
}

func main() {
	TestSelect3()
}
