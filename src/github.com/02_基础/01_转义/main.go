package main

import "fmt"

func main() {
	fmt.Println("姓名\t年龄\t籍贯\t住址\njohn\t12\t河北\t北京")
	fmt.Printf("%d %x %X %o %d \n", 18, 11, 12, 18, 18) // 18 b C 22 18
	fmt.Printf("%f \n", 1.5)                            // 1.500000
	fmt.Printf("%t \n", false)
	fmt.Printf("%c \n", 97)            // a
	fmt.Printf("%s \n", "hello world") // hello world
	fmt.Printf("%q \n", "hello world") // "hello world"
	fmt.Printf("%v \n", "你好")          // 你好
	fmt.Printf("%T \n", false)         // bool

	i := 5
	fmt.Printf("%p \n", &i)   // 0xc00001a088
	fmt.Printf("%d%% \n", 20) // 20%
}
