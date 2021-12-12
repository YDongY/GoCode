package _5_obj_cache

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a ne object.")
			return 100
		}}

	v := pool.Get().(int)
	t.Log(v) // 100

	pool.Put(3)

	// runtime.GC() GC 会清楚 sync.Pool 中的缓存对象

	v1, _ := pool.Get().(int)
	t.Log(v1) // 3，如果是 GC 了，这里结果是 100
}

func TestSyncPoolMultiGoroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a ne object.")
			return 100
		}}

	pool.Put(10)
	pool.Put(10)
	pool.Put(10)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			t.Log(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
