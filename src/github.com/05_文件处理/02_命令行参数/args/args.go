package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("命令行参数:", len(os.Args))

	// ./args aa bb
	// [./args aa bb]
	fmt.Println(os.Args)
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}
}
