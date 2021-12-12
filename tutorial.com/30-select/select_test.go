package _0_select

import (
	"testing"
	"time"
)

func TestSelect1(t *testing.T) {
	// select 可以同时监听一个或多个 channel，直到其中一个 channel ready
	output1 := make(chan string)
	output2 := make(chan string)

	go func(ch chan string) {
		time.Sleep(time.Second * 5)
		ch <- "test1"
	}(output1)

	go func(ch chan string) {
		time.Sleep(time.Second * 2)
		ch <- "test2"
	}(output2)

	select {
	case s1 := <-output1:
		t.Log("s1 =", s1)
	case s2 := <-output2:
		t.Log("s2 =", s2)
	}
}

func TestSelect2(t *testing.T) {
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
		t.Log("int:", value)
	case value := <-ch2:
		t.Log("string:", value)
	}
	t.Log("main 结束")
}

func TestSelect3(t *testing.T) {
	// 用于判断管道是否存满
	// 创建管道
	output1 := make(chan string, 10)
	// 子协程写数据
	go func() {
		for {
			select {
			// 写数据
			case output1 <- "hello":
				t.Log("write hello")
			default:
				t.Log("channel full")
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()
	// 取数据
	for s := range output1 {
		t.Log("res:", s)
		time.Sleep(time.Second)
	}
}

func TestTimeOut(t *testing.T) {
	retCh := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 1)
		retCh <- "hello"
	}()

	select {
	case ret := <-retCh:
		t.Log("ret =", ret)
	case <-time.After(time.Second * 2):
		t.Error("time out")
	}
}
