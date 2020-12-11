package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.OpenFile("text.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("error openfile")
	}
	defer file.Close()

	n, err := file.WriteString("123")
	if err != nil {
		fmt.Println("error WriteString")
	}
	fmt.Println(n)

	ret, err := file.Seek(5, io.SeekStart)
	fmt.Println(ret)

	n, err = file.WriteAt([]byte("hello world"), ret)

	n, err = file.Write([]byte("hello wolrd"))
	fmt.Println(n)

	reader := bufio.NewReader(file)
	// buf, err := reader.ReadBytes('\n')

	// buf := make([]byte, 4096)
	fmt.Println(reader)
}
