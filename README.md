# GO 入门

## Hello World

```go
package main // 特殊的包，用来定义一个独立的可执行程序

import "fmt" // 导包，必须在 package 之后

// main 是程序执行的入口
func main() { // `{` 必须和 关键在 func 同一行
	fmt.Println("hello 世界") // 原生支持 unicode
	// Println 输出
}
```

## 声明

名称开头字母或下划线，后面是任意数量字符、数字。

- 一个实体名称声明在函数中，则在函数局部有效。
- 声明在函数外，对整个包有效
- 如果声明的实体名称以**大写字母开头**，则表示其可跨包使用，例如：`fmt.Println()`，函数首字母大写

GO 语言编程风格，作用域越小，变量名越短，否则越长，且变量名采用**驼峰式**风格

## 变量

```go
var name string = "hello" // 声明+初始化

// 或者
var name string // 声明
name = "hello"  // 初始化

// 或者

name := "hello" // 初始化+自动类型检查
```

对于 GO 语言而言定义之后的变量，必须使用，如果不则会编译失败。对于不同类型的变量声明具有默认值。

对于数值类型默认值是 0 、布尔类型默认值为 flase、对于字符串类型默认值为 ""、对于接口和引用类型（slice，指针，map，channel，函数）默认值为 nil、对于像数组或结构体这样的复合类型，默认值是其所有元素的默认值。

### 一次性声明多个变量

```go
var i, j, k int                  // int,int,int
var b, f, s = true, 2.3, "hello" // bool,float64,string
a, b, c := 100, "tom", 888
```

### 变量的修改

```go
// 修改变量的值，必须为同一数据类型
var b string = "hello world"
b = "hello go"

// b :=100 // 同一作用域不能出现多个相同的变量名，不同类型也不可以
```

## 常量

```go
// 常量必须赋值
const num int = 100

// num = 200 // 常量不能修改

const num2 = num / 2 // 正确 num 必须是常量，得到的结果才是常量，否则编译器报错
```

通过 const 关键字定义常量，声明的常量必须要初始化值，而且常量的值无法修改。如果是通过计算的到的常量，则参与计算的也必须是常量。

## 数据类型

## 数据类型转换

基本数据类型相互转换,不能自动转换,必须强制转换（显式转换），格式 `type(变量名)`，而且一般都是类型小的转换类型大的，反之可能会导致结果范围溢出。

对于转换失败的类型，默认会转换为对应类型的默认值。

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {

	i := 42
	f := float64(i)　// int -> float64 转换完成之后,i 的数据类型并没有改变

	var a int64 = 1234567
	var b int8 = int8(a)
	fmt.Println("b=", b) //-121 // int64 -> int8 范围溢出


	// 基本数据类型与string类型相互转换
	var num1 int = 99
	var num2 float64 = 23.456
	var boo bool = true
	var char byte = 'h'
	var str string

	str = fmt.Sprintf("%c", char) // byte -> string
	fmt.Printf("str type %T str=%q\n", str, str)


	// 基本数据类型转 string 类型(函数转换)

	str = strconv.Itoa(num1)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatInt(int64(num1), 10)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatFloat(num2, 'f', 10, 64) //
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatBool(boo)
	fmt.Printf("str type %T str=%q\n", str, str)

	// string 类型转换基本数据类型(函数转换

	var str2 string = "true"
	var bo bool
	bo, _ = strconv.ParseBool(str2)
	fmt.Printf("bo type %T b=%v\n", bo, bo)

	var str3 string = "123456"
	var numInt64 int64
	numInt64, _ = strconv.ParseInt(str3, 10, 64)
	fmt.Printf("numInt64 type %T numInt64=%v\n", numInt64, numInt64)

	var str4 string = "123.456"
	var numFloat64 float64
	numFloat64, _ = strconv.ParseFloat(str4, 64)
	fmt.Printf("numFloat64 type %T numFloat64=%v\n", numFloat64, numFloat64)

	// go 对于不能转换成功的，会转换成默认值
	// 例:hello 转 int 类型，结果为 0
	// 例:hello 转 bool 类型，结果为 false
	var str5 string = "hello"
	// var num int64
	var num int64 = 11
	num, _ = strconv.ParseInt(str5, 10, 64)
	fmt.Printf("num type %T num=%v\n", num, num) //0

}
```

## 指针

变量是存储值的地方，而指针的值是一个变量的地址，一个指针指示值所保存的位置

```go
var i int = 10 // 定义一个 int 类型变量 i
fmt.Println("i的地址＝", &i) // 0xc000098010
j = &i
*j = 2
fmt.Println(i) // 2
```

可以通过 `&name` 来获取变量值所存储的地址，还可以通过 `*name=value`　来修改 `&name` 地址的值。

## 运算符

### 数值运算符

```go
fmt.Println(10 / 4)          // 2
fmt.Println(10 * 1.0 / 4)    // 2.5
fmt.Println(float32(10 / 4)) // 2
fmt.Println(4 / 10)          // 0
fmt.Println(4 * 1.0 / 10)    // 0.4
fmt.Println(10 % 4)          // 2
i := 1
i++
fmt.Println(i)               // 2
i--
fmt.Println(i)               // 1
```

> 注意：对于 '/' 相当于整除，除非除数是 float 类型，其次对于 GO 语言中的 ++ 和 -- 只能单独使用，不能　--i 或者 ++i 或者 j:=i++ i=i++

### 关系运算符(>,<,=,!=,>=,<=)

```go
fmt.Println(1 == 1)  // true
fmt.Println(1 != -1) // true
fmt.Println(1 < -1)  // false
fmt.Println(1 > -1)  // true
fmt.Println(1 >= -1) // true
fmt.Println(1 <= -1) // false

flag := 1 > 1
fmt.Println(flag)    // false

flag2 := 1 == 1
fmt.Println(flag2)   // true
```

### 逻辑运算符

```go
fmt.Println(1 == 1 && 2 > 1) // true
fmt.Println(1 == 1 && 2 > 3) // false
fmt.Println(1 == 1 || 2 > 3) // true
fmt.Println(1 != 1 || 2 > 3) // false

fmt.Println(!true)     // false
fmt.Println(1 != 1)    // false
fmt.Println(!(1 != 1)) // true
```

对于　&& 如果左边为假，则不判断右边，|| 如果右边为真，则不判断右边

## 输入输出

```go
package main

import "fmt"

func main() {
	var name string
	var age byte
	var sal float32
	var isPass bool
	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水：")
	fmt.Scanln(&sal)
	fmt.Printf("%v,%v,%v,%v \n", name, age, sal, isPass)
}
```

## 流程控制

### `if-else`

```go
if xxx >= 18 {
	// ...
} else {
	// ...
}
```

### `if-else if-else`

```go
var num int = 500

if num >= 100 && num < 1000 {
	// ...
} else if num >= 1000 {
	// ...
} else {
	// ...
```

### `switch` 每个 case 不用添加 break 自动在结尾添加 break

```go
var key byte
fmt.Println("请输入一个字符：a,b,c,d,e,f,g")
fmt.Scanf("%c", &key)

switch key { // key 常量、变量、返回值
	// case 类型必须与 switch 一致，case 表达式为常量，不能重复
	case 'a':
		fmt.Println("周一")
	case 'b':
		fmt.Println("周二")
	case 'c':
		fmt.Println("周三")
	case 'd':
		fmt.Println("周四")
	case 'e':
		fmt.Println("周五")
	case 'f':
		fmt.Println("周六")
	case 'g':
		fmt.Println("周日")
	case 'h', 'j', 'k': // 可以有多个
		fmt.Println("其他")
	default:
		fmt.Println("输入有误")
	}
```

### `switch fallthrough`

```go
switch { // swithch 后面可以不写，类似于 if else 分支
case key == 'a':
	fmt.Println("a")
	fallthrough // 穿透，会接着执行后面的一个 case 并且不用判断条件是否成立
case key == 'b':
	fmt.Println("b")
	// fallthrough
	// case key == 'c':
	// 	fmt.Println("c")
}
```

### `for` 循环

```go
for i := 1; i <= 10; i++ {
	fmt.Println("hello world")
}
```

### `for` 死循环

```go
flag := 1
for { // 不加条件
	fmt.Println(flag)
	flag++
	if flag > 100 {
		break
	}
}
```

### `for` 实现 while 循环

```go
flag2 := 100
for flag2 >= 0 {
	fmt.Println(flag2)
	flag2--
}
```

### `goto`

```go
fmt.Println("hello world1")
goto lable1 // 执行到该位置，跳转到 label1 位置接着执行
fmt.Println("hello world2")
fmt.Println("hello world3")
lable1:
fmt.Println("hello world4")
fmt.Println("hello world5")
fmt.Println("hello world6")
```

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

## 字符串操作

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// 1. 字符串长度
	str := "hello 世界"
	fmt.Println(len(str)) // 12 ,ascii 字符占用 1 字节，汉子占用 3 个字节

	// 2. 字符串遍历
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i]) // hello ä¸ç
	}
	fmt.Println()

	// 转换 防止乱码
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c", r[i]) // hello 世界
	}
	fmt.Println()

	// 3. 字符串转整数
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println(n)
	}

	// 4. 整数转字符串
	nStr := strconv.Itoa(123456)
	fmt.Printf("%T,%s \n", nStr, nStr) //string,123456

	// 5. 字符串转 []byte
	var bytes = []byte("hello")
	fmt.Println(bytes) //[104 101 108 108 111]

	// 6. 把 []byte 转　字符串
	bStr := string([]byte{97, 98, 99})
	fmt.Println(bStr) //abc

	// 7. 十进制转 二、八、十六

	rStr := strconv.FormatInt(123, 2)
	fmt.Println("123对应二进制：", rStr) // 1111011
	rStr = strconv.FormatInt(123, 8)
	fmt.Println("123对应八进制：", rStr) // 173
	rStr = strconv.FormatInt(123, 16)
	fmt.Println("123对应十六进制：", rStr) // 7b

	// 8. 查找子串
	fmt.Println(strings.Contains("hello world", "hello")) // true
	fmt.Println(strings.Contains("hello world", "lol"))   // false

	// 9. 返回不重复子串
	fmt.Println(strings.Count("ceheses", "e")) // 3
	fmt.Println(strings.Count("ceheses", "a")) // 0

	// 10. 字符串比较(==,区分大小写),不区分大小写如下
	fmt.Println(strings.EqualFold("abc", "ABc")) // tue
	fmt.Println("abc" == "ABC")                  // false

	// 11. 返回子串第一次出现的索引
	fmt.Println(strings.Index("abc_hello", "hello"))     // 4
	fmt.Println(strings.Index("abc_hello", "z"))         // -1
	fmt.Println(strings.LastIndex("abc abc abc", "abc")) // 8

	// 12. 子串替换，返回新串，1　表示替换第一个，-1 表示替换所有
	fmt.Println(strings.Replace("abc abc", "abc", "xyz", 1))  // xyz abc
	fmt.Println(strings.Replace("abc abc", "abc", "xyz", -1)) // xyz xyz

	// 13. 字符串拆分字符数组
	fmt.Println(strings.Split("hello,world,go", ",")) // [hello world go]

	// 14. 大小写转换
	fmt.Println(strings.ToUpper("Hello")) // HELLO
	fmt.Println(strings.ToLower("HELLO")) // hello

	// 15. 去掉左右两边空格
	fmt.Printf("%q\n", strings.TrimSpace(" hello world   ")) //"hello world"

	// 16. 指定去掉两边空格
	fmt.Println(strings.Trim(" ! hel!lo! ", " !")) //hel!lo

	// 17. 前缀，后缀
	fmt.Println(strings.HasPrefix("helloworld", "he")) // true
	fmt.Println(strings.HasSuffix("helloworld", "d"))  // true

}
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

## 错误处理

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

## 结构体

```go
type Cat struct {
	Name  string // 名称首字母大写可以再其他包使用
	Age   int
	Color string
}

// 方式一
var cat Cat
cat.Name = "小白"
cat.Age = 3
cat.Color = "白"

// 方式二

cat := Cat{
	Name: "小白",
	Age:  3,
	Color: "白",
}

// 方式三
cat := ＆Cat{
	Name: "小白",
	Age:  3,
	Color: "白",
}

　　　　　　默认值

　　　　　　　 0x000191
		------------                                ------------
cat---> |    ""    | Name           cat.Name="小白"  |   "小白" |
		------------                                ------------
		|    0     | Age
		------------
		|    ""    | Color
		------------


// ---------- 注意：结构体属于值类型　----------------
var cat1 Cat
cat1.Name = "小白"
cat2 := cat1      // 这里相当于值拷贝，cat2 的修改不会影响cat1

cat2.Age = 3
cat2.Name = "小黑"

fmt.Println(cat2) // {小黑 3 }
fmt.Println(cat1) // {小白 0 }

// 所以 cat2 的修改不会影响 cat1

// 但是如果想要达到修改的目的，可以通过指针进行修改

cat3 := &cat1

(*cat3).Name = "小花" // 等价 cat3.Name = "小花"

fmt.Println("修改后的:", cat1) // {小花 0 } cat1 的字段值发生了变化
```

## 方法

```go
package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) test() { // 給 Person 类绑定一个方法，p　变量可以随意，属于值拷贝
	p.Name = "jack"     // 此处修改不会改变原 Person 的Name
	fmt.Println(p.Name) // jack
}

func (p Person) func1() { // 无参无返回值
	// ...
}

func (p Person) func2(n int) { // 有参无返回值
	// ...
}

func (p Person) func3(n int, m int) int { // 有参有返回值
	return // ...
}

func (p *Person) func4(name string){ // 接收 Person结构体地址，内部的字段的修改会改变原 Person 的值
	// (*p).Name = name　等价于
	p.Name = name
}

func (p *Person) String() string { // 类似于 Python 中的　__str__，方法名称，方法必须是 String
	str := fmt.Sprintf("Person<Name:%s>", p.Name)
	return str
}

func main() {
	// ...
}
```

## 继承

```go
package main

import "fmt"

// ------------------- 公共　-----------------------
// 公共
type Student struct {
	Name  string
	Age   int
	Score int
}

// 公共 方法
func (stu *Student) func1() {
	// ...
}

func (stu *Student) func2(score int) {
	// ...
}


// ------------------- 小学生 ----------------------

type Pupil struct {
	Student // 继承
}

func (p *Pupil) testing() {
	// ...
}

// ------------------ 大学生 -----------------------

type Graduate struct {
	Student // 继承
}

func (g *Graduate) testing() {
	// ...
}
```

Graduate，Pupil　结构体内部都声明了 Student，就表示当前两个结构体继承了 Student，他们可以使用 Student 结构的字段和方法。如果存在与 Student 相同的方法，则会调用自己的。同理结构体字段也会优先查找自己的结构体。

如果想要调用 Student (即继承过来的结构体定义的方法，本身也存在同名方法)方法，可以使用**显式调用**

```go
type A struct {
	Name string
}

type B struct {
	Name string
}

type C struct {
	A
	B
}

var c C
c.Name = "tom" // error
c.A.Name = "tom" // ok 显式调用继承的字段，同理方法也是如此
c.B.Name = "tom" // ok
```

### 有名结构体

```go
type A struct {
	Name string
}

type D struct {
	a A // 有名结构体
}

// 有名结构体
var d D
d.a.Name = "tom" // 调用时不能通过 A 调用，类似于起了一个别名
```

### 结构体赋值

```go

type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

// 结构体赋值
tv := TV{
	Goods{
		Name:  "电视机",
		Price: 5000,
	},
	Brand{
		Name:    "海尔",
		Address: "北京",
	},
}

fmt.Println(tv) // {{电视机 5000} {海尔 北京}}

tv2 := TV2{
	&Goods{
		Name:  "电视机",
		Price: 5000,
	},
	&Brand{
		Name:    "海尔",
		Address: "北京",
	},
}
fmt.Println(tv2)                    // {0xc000114000 0xc000114020}
fmt.Println(*tv2.Goods, *tv2.Brand) // {电视机 5000} {海尔 北京}
```

## 接口

```go
package main


// 声明一个接口
type Usb interface {
	// 声明两个方法
	Start()
	Stop()
}

type Phone struct {} // 定义一个 Phone 结构体为空

func (p Phone) Start() {　// Phone 实现 Usb 接口的 Start方法
	fmt.Println("手机开始工作")
}

func (p Phone) Stop() {　// Phone 实现 Usb　接口的 Stop方法
	fmt.Println("手机停止工作")
}

type Camera struct {} // 同上

func (c Camera) Start() {
	fmt.Println("相机开始工作")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

type Computer struct {}

func (c Computer) Working(usb Usb) { // 实现了 Usb 接口
	usb.Start()
	usb.Stop()
}

func main() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	computer.Working(phone)
	computer.Working(camera)
｝
```

GO 语言中的接口是基于方法的，只要一个结构体实现了某个接口中定义的所有方法，就表明当前结构体实现了该接口

## 类型断言

```go


```


## IO

### `read`

```go
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// -------- 一次性读入，不适合大文件，也不需要打开和关闭文件 ------------
	content, err := ioutil.ReadFile("../test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v \n", content) // []byte
	fmt.Println(string(content))


	// ---------- 带有缓冲区读取 --------------------------------------
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	fmt.Printf("file=%v \n", file)

	defer file.Close() // 该函数结束时关闭文件对象

	// 创建带缓冲的 Reader
	reader := bufio.NewReader(file)
	// 循环读取文件内容
	for {
		str, err := reader.ReadString('\n') // 读取到换行结束
		if err == io.EOF {                  // 表示文件末尾
			break
		}
		fmt.Printf(str)
	}
}
```

### `write`

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func write() {
	// 创建文件并写入
	f, err := os.OpenFile("notes.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err=", err)
	}

	// 关闭
	defer f.Close()

	str := "hello go\r\n"

	// 带缓冲写入
	writer := bufio.NewWriter(f)

	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	// 将缓冲 flush 到磁盘
	writer.Flush()
}

func coverWrite() {
	// 清空源文件，在写入
	f, err := os.OpenFile("notes.txt", os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("open file err=", err)
	}
	defer f.Close()

	str := "hello world\r\n"

	// 带缓冲写入
	writer := bufio.NewWriter(f)

	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}

	// 将缓冲 flush 到磁盘
	writer.Flush()
}

func appendWrite(){
	// 追加写入
	f, err := os.OpenFile("notes.txt", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file err=", err)
	}
	defer f.Close()

	str := "你好，世界\r\n"

	// 带缓冲写入
	writer := bufio.NewWriter(f)

	for i := 0; i < 1; i++ {
		writer.WriteString(str)
	}

	// 将缓冲 flush 到磁盘
	writer.Flush()
}

func main() {
	// write()
	coverWrite()
	appendWrite()
}
```