package main

import (
	"fmt"
	"reflect"
)

func reflectTest(b interface{}) {
	rVal := reflect.ValueOf(b)
	// 查看 rVal kind
	fmt.Println(rVal.Kind()) // ptr

	// 设置
	// Elem 返回 v 持有的接口保管的值的 Value 封装，或者 v 持有的指针指向的值的 Value 封装
	// func (v Value) Elem() Value
	rVal.Elem().SetInt(20)
}

func main() {
	var num int = 10
	reflectTest(&num)
	fmt.Println(num) // 20
}
