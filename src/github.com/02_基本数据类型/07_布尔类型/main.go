package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//只能用true和false,不可以用0或非0替代
	var b = false
	fmt.Println("b=", b)
	fmt.Println("b　占用的空间大小:", unsafe.Sizeof(b))
}
