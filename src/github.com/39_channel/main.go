package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	// channel 声明
	var intChan chan int

	// var intChan chan<- int 只写的 channel
	// var intChan <-chan int 只读的　channel

	intChan = make(chan int, 3)
	fmt.Println(intChan) // 0xc000076080

	// channel 写入数据
	intChan <- 10
	num := 211
	intChan <- num

	// 注意写入channel的数据不能超过容量

	// channel cap
	fmt.Println(len(intChan), cap(intChan)) // 2,3

	// channel 取数据
	var num2 int
	num2 = <-intChan
	<-intChan                               // 取出数据可以丢掉
	fmt.Println(num2)                       // 10
	fmt.Println(len(intChan), cap(intChan)) // 0,3

	// 注意 channel 没有数据时，继续取数据会 deadlock

	// 存放任意类型管道
	allChan := make(chan interface{}, 3)
	allChan <- 10
	allChan <- "tom"
	cat := Cat{"小白", 4}
	allChan <- cat

	<-allChan
	<-allChan

	newChat := <-allChan

	fmt.Println(newChat) // {小白 4}
	// fmt.Println(newChat.Name) // 无法直接获取字段，需要类型断言
	fmt.Println(newChat.(Cat).Name) // 小白

	// channel 关闭，channel 关闭之后只能读不能写
	var intChan2 chan int
	intChan2 = make(chan int, 3)

	intChan2 <- 10
	intChan2 <- 20
	close(intChan2)

	// intChan2 <- 100 // // panic: send on closed channel
	fmt.Println(<-intChan2) // 10
	fmt.Println(intChan2)

	// 遍历 channel
	var intChan3 chan int
	intChan3 = make(chan int, 100)

	for i := 0; i < 100; i++ {
		intChan3 <- i
	}
	close(intChan3) // 遍历时不关闭 channel 会出现 deadlock
	// 管道返回一个值
	for v := range intChan3 {
		fmt.Println(v)
	}

}
