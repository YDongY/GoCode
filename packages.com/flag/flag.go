package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// os.Args 是一个 []string
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}

	// name := flag.String("name", "张三", "用户名")
	// age := flag.Int("age", 18, "年龄")
	// fmt.Println(*name, *age)

	var (
		username string
		age      int
	)
	flag.StringVar(&username, "name", "张三", "用户名")
	flag.IntVar(&age, "age", 18, "年龄")

	flag.Parse()

	fmt.Println(username, age)

}
