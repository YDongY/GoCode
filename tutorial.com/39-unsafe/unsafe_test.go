package _9_unsafe

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

func TestUnsafe(t *testing.T) {
	i := 10
	f := *(*float64)(unsafe.Pointer(&i)) // unsafe.Pointer() 可以转换为任意类型指针
	fmt.Println(f)                       // 5e-323
}

func TestAtomic(t *testing.T) {
	var shareBufPtr unsafe.Pointer
	writeDataFunc := func() {
		var data []int
		for i := 0; i < 10; i++ {
			data = append(data, i)
		}

		// 原子操作，把数据放在 shareBufPtr 指针上
		atomic.StorePointer(&shareBufPtr, unsafe.Pointer(&data))
	}

	readDataFunc := func() {
		// 根据 shareBufPtr 获取数据
		data := atomic.LoadPointer(&shareBufPtr)
		fmt.Println(data, *(*[]int)(data))
	}
	var wg sync.WaitGroup
	writeDataFunc()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				writeDataFunc()
				time.Sleep(time.Millisecond * 100)
			}
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				readDataFunc()
                time.Sleep(time.Millisecond * 100)
			}
			wg.Done()
		}()
	}
    wg.Wait()

}
