package main

import (
	"fmt"
)

func makeChan() {
	var ch1 chan int
	var ch2 chan bool
	var ch3 chan []int

	fmt.Println(ch1, ch2, ch3) // <nil> <nil> <nil>

	// 声明的通道后需要使用 make 函数初始化之后才能使用  make(chan 元素类型, [缓冲大小])
	// ch4 := make(chan int)
	// ch5 := make(chan bool)
	// ch6 := make(chan []int)
}

func NoBufChan() {
	/** 创建的是无缓冲的通道，无缓冲的通道只有在有人接收值的时候才能发送值 ，否则 deadlock*/
	ch := make(chan int)
	ch <- 10
	fmt.Println("发送成功")

	// fatal error: all goroutines are asleep - deadlock!
}

func BufChan() {
	/** 内置的 len 函数获取通道内元素的数量，使用 cap 函数获取通道的容量 */
	ch := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
	ch <- 10
	fmt.Println(<-ch)
	fmt.Println("发送成功")
}

func CloseChan(){
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 开启 goroutine 将 0~100 的数发送到 ch1 中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	// 开启 goroutine 从 ch1 中接收值，并将该值的平方发送到 ch2 中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值 ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()

	// 在主 goroutine 中从 ch2 中接收值打印
	for i := range ch2 { // 通道关闭后会退出 for range 循环
		fmt.Println(i)
	}
}

func counter(out chan<- int) {
	// chan<- int 是一个只能发送的通道，可以发送但是不能接收
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) {
	// <-chan int 是一个只能接收的通道，可以接收但是不能发送
	for i := range in {
		fmt.Println(i)
	}
}

func SingleChan()  {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}


func main() {
	SingleChan()
}
