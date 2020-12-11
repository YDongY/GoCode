package main

import (
	"fmt"
	"sync"
)

func main() {
	var scene sync.Map

	// 存值
	scene.Store("name","jack")
	scene.Store("age",20)
	scene.Store("score",99.9)

	// 取值
	fmt.Println(scene.Load("name")) // jack true

	// 删除键值对
	scene.Delete("score")

	// 遍历
	scene.Range(func(k,v interface{}) bool{
		fmt.Println(k,v)
		return true
	})
}
