package main

import (
	"fmt"
	cf "golang-dev/tutorial.com/13-import-package/conf" // 包的别名
	"golang-dev/tutorial.com/13-import-package/utils"
)

func main() {
	fmt.Println("main")
	fmt.Println(utils.Name)
	fmt.Println(cf.Num)
}
