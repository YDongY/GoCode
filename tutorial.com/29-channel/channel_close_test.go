package _9_channel

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan<- int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		// ch <- 1 // panic: send on closed channel
		wg.Done()
	}()
}

func dataConsumer(ch <-chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		fmt.Println(<-ch) // 关闭通道后继续获取，会立即返回通道类型的零值
		wg.Done()
	}()
}

func TestChannelClose(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataConsumer(ch, &wg)
	wg.Wait()
}
