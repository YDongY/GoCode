package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// 1. 字符串长度
	str := "hello 世界"
	fmt.Println(len(str)) // 12 ,ascii 字符占用 1 字节，汉子占用三个字节

	// 2. 字符串遍历
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i]) // hello ä¸ç
	}
	fmt.Println()

	// 按 Unicode字符遍历字符串
	for _,s := range str{
		fmt.Printf("%c %d \n", s,s)
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
