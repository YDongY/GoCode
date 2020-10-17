package main

import "fmt"

func main() {

	//字符串拼接
	str := "hello" + "world"
	fmt.Println("str = ", str)

	//当字符串，可换行,+号必须在结尾
	str2 := "hello" + "world" + "go" +
		"python" + "c" + "java" + "php" +
		"vb" + "js" + "c#" + "object-c"
	fmt.Println("str2 = ", str2)
}
