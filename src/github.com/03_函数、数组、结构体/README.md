## 函数

```go
func test(i interface{}){ // 接收任意类型参数
	// ...
}

func test1() int { // 无参有返回值
	return 100
}

func test2(a, b int)　{　// 有参无返回值
	// ...
}

func test3(a int, b int) (int, int) { // 有参有返回值
	return b, a
}

func test4(a int, b int) (x int, y int) { // 返回值命名
	x = b
	y = a
	return // 直接返回 x,y
}

func test5(n *int) { // 传递地址
	fmt.Println(n)  // 0xc0000160e0
	fmt.Println(&n) // 0xc00000e030
	fmt.Println(*n) // 100
	*n += 1
	fmt.Println(*n) // 101
}

func test6(funcvar func() string) string { // 传递函数
	return funcvar()
}

func test7(n1 int, args ...int) int { // 函数接收可变参数
	// ...
}

res := func(n1 int, n2 int) int {  // 使用匿名函数，求和
	return n1 + n2
}(10, 20)


res2 := func(n1 int, n2 int) int { // 匿名函数赋值给变量
	return n1 + n2
}

fmt.Println(res2(10, 20))


func test8() func(int) int { // 闭包
	var n int = 10
	return func(x int) int {
		n += x
		return n
	}
}

fmt.Println(addUpper()(2))
s := addUpper()

fmt.Println(s(1)) // 11
fmt.Println(s(2)) // 13
//一次调用之后，下次调用变量 n 还是上次一调用之后的值
```


## 时间日期

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%T,%v \n", now, now) // time.Time,2020-10-17 22:44:54.647712027 +0800 CST m=+0.000046700

	// 获取日期
	fmt.Println("年：", now.Year())
	fmt.Println("月：", now.Month())
	fmt.Println("月：", int(now.Month()))
	fmt.Println("日：", now.Day())
	fmt.Println("时：", now.Hour())
	fmt.Println("分：", now.Minute())
	fmt.Println("秒：", now.Second())

	// 格式化

	fmt.Println(now.Format("2006/01/02 15:04:05")) // 2020/10/17 22:49:20
	fmt.Println(now.Format("2006-01-02 15:04:05")) // 2020-10-17 22:49:37
	fmt.Println(now.Format("2006-01-02"))          // 2020-10-17
	fmt.Println(now.Format("01"))                  // 10
	fmt.Println(now.Format("2006"))                // 2020

	// 时间常量
	// const (
	// 	Nanosecond  Duration = 1
	// 	Microsecond          = 1000 * Nanosecond
	// 	Millisecond          = 1000 * Microsecond
	// 	Second               = 1000 * Millisecond
	// 	Minute               = 60 * Second
	// 	Hour                 = 60 * Minute
	// )
	// 随机数字
	// unix时间戳=1602948366, unixnaon时间戳=1602948366056742753
	fmt.Printf("unix时间戳=%v, unixnaon时间戳=%v", now.Unix(), now.UnixNano())
}
```

## 数组

1. 数组数据元素类型一致，长度固定，不能动态增长
2. var arr[]int ，arr 是切片，不是数组
3. 数组中的元素可以是任意类型
4. 数组初始，数值类型默认 0 ,字符串类型默认"",bool 类型默认 false
5. 数组是值类型，函数传递会进行值拷贝

```go
var arr [6]float64
fmt.Printf("%p \n", &arr)    //0xc000018420 ,数组地址就是第一个元素地址
fmt.Printf("%p \n", &arr[0]) // 0xc000018420
fmt.Printf("%p \n", &arr[1]) // 0xc000018428 第一个元素地址加上 8 字节，即加上数据类型的大小
```

### 数组初始化方式

```go
var arr1 [3]int = [3]int{1, 2, 3}
fmt.Println(arr1)

var arr2 = [3]int{1, 2, 3}
fmt.Println(arr2)

var arr3 = [...]int{1, 2, 3}
fmt.Println(arr3)

var arr4 = [...]int{1: 800, 0: 900, 2: 999} // 指定下标元素
fmt.Println(arr4)                           // [900 800 999]

arr5 := [...]string{"tom", "jack", "mike"} // [mike tom jack]
```

### 数组遍历 for-range

```go
for index, val := range arr5 {
	fmt.Println(index, val)
}
```

## 二维数组

```go
var arr1 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
fmt.Println(arr1) // [[1 2 3] [4 5 6]]

var arr2 [2][3]int = [...][3]int{{1, 2, 3}, {4, 5, 6}}
fmt.Println(arr2) // [[1 2 3] [4 5 6]]

var arr3 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
fmt.Println(arr3) // [[1 2 3] [4 5 6]]

var arr4 = [2][3]string{{"a", "b", "c"}, {"d", "e", "f"}}
fmt.Println(arr4) // [[a b c] [d e f]]

arr5 := [2][3]string{{"a", "b", "c"}, {"d", "e", "f"}}
fmt.Println(arr5) // [[a b c] [d e f]]
```

## 切片

```go
var intArr [5]int = [...]int{1, 2, 3, 4, 5} // 定义数组

slice := intArr[1:3]    // slice 引用　intArr 数组,其实下标 1 - 3，不包含3
fmt.Println(slice)
fmt.Println(len(slice)) // 2 长度
fmt.Println(cap(slice)) // 4 容量


// 数组
------------------------------------------
| 1  |   2   |    3  |     4   |    5    |
------------------------------------------
			↑(地址)
------------------------------------------
|  第一个元素地址 |   len    |     cap     |　　slice[1:3] = [2 3]
------------------------------------------

// slice 底层就是一个 struct 结构

type slice struct{
	ptr *[2]int
	len int // 长度
	cap int // 容量，可动态增长
}

// 切片引用了数组指定索引范围，即切片数据的修改会影响数组的元素

slice[1] = 999 // 修改切片下标为 1 的元素
fmt.Println(slice)  // [2 999]
fmt.Println(intArr) // [1 2 999 4 5]
```

### 切片的定义

```go
var arr [5]int = [...]int{1, 2, 3, 4, 5}

// 方式一，不支持负索引
var slice1 = arr[1:3]
fmt.Println(slice3[:])  // [1 2 3 4 5]
fmt.Println(slice3[3:]) // [4 5]

// 方式二　make([]type,len,cap) cap>=len
var slice2 []int = make([]int, 4, 10) // 类型，大小，容量
// slice2 [0 0 0 0]

// 方式三，注意中括号内没有指定大小
var slice3 []string = []string{"1", "2", "3", "4", "5"}
```

### 切片数据添加 append

> append 原理：先创建一个数组，源切片引用的数组元素拷贝到新数组，然后切片引用新数组，原数组会被垃圾回收

```go
var slice4 []int = []int{1, 2, 3}

// 使用 append 关键字
newSlice := append(slice4, 400, 500, 600)
fmt.Println(newSlice) // [1 2 3 400 500 600]

var slice5 []int = []int{100, 200, 300}

// 切片追加切片，注意第二个参数后面的...
slice4 = append(slice4, slice5...)　
fmt.Println(slice4) // [1 2 3 100 200 300]
```

### 切片拷贝 copy

copy 切片，数据不会相互影响，但是只有切片才能 copy

```go
// 切片 copy
var slice6 []int = []int{1, 2, 3, 4, 5}
var slice7 = make([]int, 10)
copy(slice7, slice6)
fmt.Println(slice7) // [1 2 3 4 5 0 0 0 0 0]
fmt.Println(slice6) // [1 2 3 4 5]
```

## map

```go
var map1 map[string]string
fmt.Println(map1) // map[]
// 使用 map 之前需要 make 分配数据空间
map1 = make(map[string]string, 10) // 最多存放10
map1["name"] = "ydongy"
map1["gender"] = "male"
fmt.Println(map1) // map[gender:male name:ydongy]
```

注意：在使用 map 之前必须进行 make　，　否则会报错(`panic: assignment to entry in nil map`)

### 定义 map 的方式

```go
// 方式一
var map2 map[string]string
map2 = make(map[string]string, 10)

// 方式二
var map3 = make(map[string]string, 10)

// 方式三
map4 := make(map[string]string)

map4["name"] = "tom"
map4["city"] = "北京"
map4["gender"] = "male"
fmt.Println(map4) // map[city:北京 gender:male name:tom]

// 方式四
map5 := map[string]string{
	"name": "jack",
	"age":  "20",
}

fmt.Println(map5) // map[age:20 name:jack]
```

### map 基本操作

```go
map1 := map[string]string{
	"name": "jack",
	"age":  "20",
}

delete(map1, "name") // 删除 key

map1 = make(map[string]string) // 清空map　重新创建空间

for k, v := range map { // 遍历 map
	fmt.Println(k, v)
}

var mapSlice []map[string]string // 定义 map 类型 切片，有点类似于 Python 列表嵌套字典
mapSlice = make([]map[string]string, 2) // 切片大小，2个map

newMap := map[string]string{
	"name": "ydony",
	"age":  "23",
}
mapSlice = append(mapSlice, newMap) // append 添加 map 切片
```