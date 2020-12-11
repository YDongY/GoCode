package main

import "fmt"

type Invoker interface {
	Call(interface{})
}

type Struct struct {
}

func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

// 函数体实现接口
// 函数定义为类型
type FuncCaller func(interface{})

func (f FuncCaller) Call(p interface{}) {
	f(p)
}

func main() {
	s := new(Struct)
	s.Call("hello")

	FuncCaller(func(v interface{}) {
		fmt.Println("from function", v)
	}).Call("hello")

}
