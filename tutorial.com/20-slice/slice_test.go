package slice

import (
	"testing"
)

func TestSlice(t *testing.T) {
	// 切片
	var s []int
	t.Log(s)
}

func TestCreateSliceByArray(t *testing.T) {
	var intArr [5]int = [...]int{1, 2, 3, 4, 5}
	slice := intArr[1:3]                 // slice 引用　intArr 数组,其实下标 1 - 3，不包含3
	t.Log(slice, len(slice), cap(slice)) // [2 3] 2 4

	/**
	  	      ------------------------------------------
	  	      | 1  |   2   |    3  |     4   |    5    |
	  	      ------------------------------------------
	                    ↑(地址)
	  	      ------------------------------------------
	  	      |  第一个元素地址 |   len    |     cap     |　　slice[1:3] = [2 3]
	  	      ------------------------------------------

	  	      slice 底层就是一个 struct 结构
	  	      type slice struct{
	  	  		ptr *[2]int
	  	  		len int
	  	  		cap int
	  	      }

	*/

	slice[1] = 999
	t.Log(slice)  // [2 999]
	t.Log(intArr) // [1 2 999 4 5]
}

func TestSliceInit(t *testing.T) {
	var arr [5]int = [...]int{1, 2, 3, 4, 5}

	// 方式一 通过数组获取
	var slice1 = arr[1:3]
	t.Log(slice1, len(slice1), cap(slice1)) // [2 3] 2 4

	// 方式二　make([]type,len,cap) cap>=len
	var slice2 []int = make([]int, 4, 10)   // 类型，长度，容量
	t.Log(slice2, len(slice2), cap(slice2)) // [0 0 0 0] 4 10

	// 方式三 初始化定义
	var slice3 []string = []string{"1", "2", "3", "4", "5"}
	t.Log(slice3, len(slice3), cap(slice3)) // [1 2 3 4 5] 5 5
}

func TestSliceMethod(t *testing.T) {
	// ----------- append 向切片添加元素，返回新的 切片 -----------
	var slice []int = []int{1, 2, 3}
	newSlice := append(slice, 400, 500, 600)
	t.Log(newSlice) // [1 2 3 400 500 600]

	// ----------- append 切片追加切片 -----------
	slice = append(slice, []int{100, 200, 300}...)
	t.Log(slice) // [1 2 3 100 200 300]

	// append 原理: 先创建一个数组，源切片引用的数组元素拷贝到新数组，然后切片引用新数组，原数组会被垃圾回收

	// ---------- copy 切片拷贝 ---------------
	var sortCopySlice = make([]int, 3)
	copy(sortCopySlice, []int{1, 2, 3, 4, 5})
	t.Log(sortCopySlice) // [1 2 3]

	var lengthCopySlice = make([]int, 10)
	copy(lengthCopySlice, []int{1, 2, 3, 4, 5})
	t.Log(lengthCopySlice) // [1 2 3 4 5 0 0 0 0 0]
}

func TestSliceCompare(t *testing.T) {
	// a := []int{1, 2, 3, 4}
	// b := []int{1, 2, 3, 4, 5}
	// t.Log(a == b)

	// 切片比较：invalid operation: a == b (slice can only be compared to nil)
}
