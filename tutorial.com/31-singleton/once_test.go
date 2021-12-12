package _1_singleton

import (
	"fmt"
	"sync"
	"testing"
)

type Singleton struct{}

var (
	singleInstance *Singleton
	once           sync.Once
	wg             sync.WaitGroup
)

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestGetSingletonObj(t *testing.T) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("%p\n", obj)
			wg.Done()
		}()
	}

	wg.Wait()
}
