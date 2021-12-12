package _7_goroutine

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second)
	t.Logf("counter=%d\n", counter)
}

func TestCounterGoroutineSafe(t *testing.T) {
	var mux sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer mux.Unlock()
			mux.Lock()
			counter++
		}()
	}
	time.Sleep(time.Second)
	t.Logf("counter=%d\n", counter)
}

func TestCounterWithGroup(t *testing.T) {
	var mux sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer mux.Unlock()
			mux.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait() // 等待协程执行完毕
	t.Logf("counter=%d\n", counter)
}
