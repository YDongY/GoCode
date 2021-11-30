package main

import (
	"fmt"
	cf "golab/example.com/13-import-package/conf" // cf 包的别名
	"golab/example.com/13-import-package/utils"
)

func main() {
	fmt.Println("main")
	fmt.Println(utils.HeroName)
	fmt.Println(cf.Num)
}
