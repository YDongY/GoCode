package _3_all_response

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(1 * time.Second)
	return fmt.Sprintf("the result is from %d", id)
}

func AllResponse() []string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	var resSlice []string
	for i := 0; i < numOfRunner; i++ {
		resSlice = append(resSlice, <-ch+", ")
	}
	return resSlice
}

func TestAllResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())
}
