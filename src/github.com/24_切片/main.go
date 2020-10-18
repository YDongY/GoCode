package main

import "fmt"

func main() {

	// 切片
	var s []int
	fmt.Println(s)
	// 数组
	var intArr [5]int = [...]int{1, 2, 3, 4, 5}

	slice := intArr[1:3] // slice 引用　intArr 数组,其实下标 1 - 3，不包含3
	fmt.Println(slice)
	fmt.Println(len(slice)) // 2 长度
	fmt.Println(cap(slice)) // 4 容量

	// 数组
	// ------------------------------------------
	// | 1  |   2   |    3  |     4   |    5    |
	// ------------------------------------------
	//			↑(地址)
	// ------------------------------------------
	// |  第一个元素地址 |   len    |     cap     |　　slice[1:3] = [2 3]
	// ------------------------------------------

	// slice 底层就是一个 struct 结构
	// type slice struct{
	// 	ptr *[2]int
	// 	len int
	// 	cap int
	// }

	slice[1] = 999
	fmt.Println(slice)  // [2 999]
	fmt.Println(intArr) // [1 2 999 4 5]

	// 切片使用方式
	var arr [5]int = [...]int{1, 2, 3, 4, 5}

	// 方式一
	var slice1 = arr[1:3]
	fmt.Println(slice1)

	// 方式二　make([]type,len,cap) cap>=len
	var slice2 []int = make([]int, 4, 10) // 类型，大小，容量

	fmt.Println(slice2) // [0 0 0 0]

	// 方式三
	var slice3 []string = []string{"1", "2", "3", "4", "5"}

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	// 不支持负索引
	fmt.Println(slice3[:])  // [1 2 3 4 5]
	fmt.Println(slice3[3:]) // [4 5]

	// append，返回新的 切片

	var slice4 []int = []int{1, 2, 3}

	newSlice := append(slice4, 400, 500, 600)
	fmt.Println(newSlice) // [1 2 3 400 500 600]

	var slice5 []int = []int{100, 200, 300}

	// 切片追加切片
	slice4 = append(slice4, slice5...)
	fmt.Println(slice4) // [1 2 3 100 200 300]

	// append 原理
	// 先创建一个数组，源切片引用的数组元素拷贝到新数组，然后切片引用新数组，原数组会被垃圾回收

	// 切片 copy
	var slice6 []int = []int{1, 2, 3, 4, 5}
	var slice7 = make([]int, 10)
	copy(slice7, slice6)
	fmt.Println(slice7) // [1 2 3 4 5 0 0 0 0 0]
	fmt.Println(slice6) // [1 2 3 4 5]

	// copy 切片数据不会相互影响，且只有切片才能 copy
}
