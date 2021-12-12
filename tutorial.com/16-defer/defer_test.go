package _defer

import (
	"fmt"
	"testing"
)

func Defer(s string) string {
	// 当执行到 defer 时，暂时不执行，会将 defer 后面的语句压入到栈中
	// 当函数执行完毕，再从栈中依次弹出执行
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	fmt.Println("Defer Func")
	return s
}

func TestDefer(t *testing.T) {
	t.Log(Defer("Defer"))

	/**
	  输出：
	      Defer Func
	      defer2
	      defer1
	*/
}

func DeferWithPanic(s string) string {
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	panic("这里出现了 panic") // defer 仍然执行
	fmt.Println("DeferWithPanic Func")
	return s
}

func TestDeferWithPanic(t *testing.T) {
	t.Log(DeferWithPanic("DeferWithPanic"))

	/**
	  输出：
	      defer2
	      defer1
	      panic: 这里出现了 panic
	*/
}
