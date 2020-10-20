package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file1Path := "source.txt"
	file2Path := "dest.txt"

	content, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Printf("read file err=%v\n", err)
		return
	}

	err = ioutil.WriteFile(file2Path, content, 0666)
	if err != nil {
		fmt.Printf("write file err=%v\n", err)
	}
}
