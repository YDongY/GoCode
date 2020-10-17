package main

import "fmt"

func main() {
	var name string
	var age byte
	var sal float32
	var isPass bool
	// fmt.Println("请输入姓名：")
	// fmt.Scanln(&name)
	// fmt.Println("请输入年龄：")
	// fmt.Scanln(&age)
	// fmt.Println("请输入薪水：")
	// fmt.Scanln(&sal)
	// fmt.Println("是否通过(true/false)：")
	// fmt.Scanln(&isPass)

	// fmt.Printf("%v,%v,%v,%v \n", name, age, sal, isPass)

	fmt.Println("请输入姓名，年龄，薪水，是否通过;使用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
}
