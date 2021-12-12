package array

import (
	"testing"
)

func TestArray(t *testing.T) {
	// -------------- 数组是值类型，存储多个同一类型数据 --------------

	var arr [6]float64
	t.Logf("%p \n", &arr)    // 0xc000018420 ,数组地址就是第一个元素地址
	t.Logf("%p \n", &arr[0]) // 0xc000018420
	t.Logf("%p \n", &arr[1]) // 0xc000018428 第一个元素地址加上 8 字节

	for i := 0; i < len(arr); i++ {
		arr[i] = float64(i)
	}
	t.Log(arr)
}

func TestArrayInit(t *testing.T) {
	// -------------- 数组初始化方式 --------------
	var arr1 [3]int = [3]int{1, 2, 3}
	t.Log(arr1)

	var arr2 = [3]int{1, 2, 3}
	t.Log(arr2)

	var arr3 = [...]int{1, 2, 3} // 自动推断数组元素个数
	t.Log(arr3)

	var arr4 = [...]int{1: 800, 0: 900, 2: 999} // 指定下标元素
	t.Log(arr4)                                 // [900 800 999]

	arr5 := [...]string{"tom", "jack", "mike"} // [mike tom jack]
	t.Log(arr5)
}

func TestArrayTravel(t *testing.T) {
	// -------------- for-range 遍历 --------------
	arr := [...]string{"tom", "jack", "mike"}
	for index, val := range arr {
		t.Log(index, val)
	}

	// -------------- for-i 遍历 --------------
	var byteArr [26]byte
	for i := 0; i < 26; i++ {
		byteArr[i] = 'A' + byte(i)
	}
	t.Logf("%c \n", byteArr)
}

func TestArraySection(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}
	t.Log(a[:])
	t.Log(a[2:])
	t.Log(a[:5])
}

func TestArrayCompare(t *testing.T) {
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2}
	t.Log(a == b) // false

	c := [...]int{1, 2, 3}
	t.Log(a == c) // true

	// 数组比较：必须是类型相同，所谓类型包括 [3]int 整体，而不是元素类型

	// d := [3]float64{1.1, 2.1, 3.1}
	// t.Log(a == d)
}

func TestMultidimensionalArray(t *testing.T) {
	var arr [4][6]int
	t.Log(arr) // [[0 0 0 0 0 0] [0 0 0 0 0 0] [0 0 0 0 0 0] [0 0 0 0 0 0]]

	// 定义方式
	var arr1 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	t.Log(arr1) // [[1 2 3] [4 5 6]]

	var arr2 [2][3]int = [...][3]int{{1, 2, 3}, {4, 5, 6}}
	t.Log(arr2) // [[1 2 3] [4 5 6]]

	var arr3 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	t.Log(arr3) // [[1 2 3] [4 5 6]]

	var arr4 = [2][3]string{{"a", "b", "c"}, {"d", "e", "f"}}
	t.Log(arr4) // [[a b c] [d e f]]

	arr5 := [2][3]string{{"a", "b", "c"}, {"d", "e", "f"}}
	t.Log(arr5) // [[a b c] [d e f]]
}

// -------------- 数组细节 --------------
// 1. 数组数据元素类型一致，长度固定，不能动态增长
// 2. var arr[]int ，arr 是切片，不是数组
// 3. 数组中的元素可以是任意类型
// 4. 数组初始，数值类型默认 0 ,字符串类型默认 "",bool 类型默认 false
// 5. 数组是值类型，函数传递会进行值拷贝
