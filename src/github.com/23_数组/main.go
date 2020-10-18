package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 数组是值类型，存储多个同一类型数据

	var arr [6]float64
	fmt.Printf("%p \n", &arr)    //0xc000018420 ,数组地址就是第一个元素地址
	fmt.Printf("%p \n", &arr[0]) // 0xc000018420
	fmt.Printf("%p \n", &arr[1]) // 0xc000018428 第一个元素地址加上 8 字节

	for i := 0; i < len(arr); i++ {
		arr[i] = float64(i)
	}
	fmt.Println(arr)

	// 数组初始化方式
	var arr1 [3]int = [3]int{1, 2, 3}
	fmt.Println(arr1)

	var arr2 = [3]int{1, 2, 3}
	fmt.Println(arr2)

	var arr3 = [...]int{1, 2, 3}
	fmt.Println(arr3)

	var arr4 = [...]int{1: 800, 0: 900, 2: 999} // 指定下标元素
	fmt.Println(arr4)                           // [900 800 999]

	arr5 := [...]string{"tom", "jack", "mike"} // [mike tom jack]
	fmt.Println(arr5)

	// for-range 遍历

	for index, val := range arr5 {
		fmt.Println(index, val)
	}

	// 数组细节
	// 1. 数组数据元素类型一致，长度固定，不能动态增长
	// 2. var arr[]int ，arr 是切片，不是数组
	// 3. 数组中的元素可以是任意类型
	// 4. 数组初始，数值类型默认 0 ,字符串类型默认"",bool类型默认 false
	// 5. 数组是值类型，函数传递会进行值拷贝

	//
	var byteArr [26]byte

	for i := 0; i < 26; i++ {
		byteArr[i] = 'A' + byte(i)
	}

	fmt.Printf("%c \n", byteArr)

	// 求数组最大值，返回下标

	var intArr [6]int = [...]int{1, -1, 9, 90, 11, 9000}
	max := intArr[0]
	maxIndex := 0
	for index, val := range intArr {
		if val > max {
			max = val
			maxIndex = index
		}
	}
	fmt.Println(maxIndex)

	// 随机数

	var intArr2 [6]int
	// 为了生成的随机数不一样，我们需要给一个seed值
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(intArr2); i++ {
		intArr2[i] = rand.Intn(100)
	}
	fmt.Println(intArr2)

	temp := 0
	for i := 0; i < len(intArr2)/2; i++ {
		temp = intArr2[len(intArr2)-1-i]
		intArr2[len(intArr2)-1-i] = intArr2[i]
		intArr2[i] = temp
	}
	fmt.Println(intArr2)
}
