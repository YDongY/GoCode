package main

import (
	"fmt"
	"os"
)

func testPrintln() {
	fmt.Println("testPrintln")
}

func testPrint() {
	fmt.Print("testPrint")
}

func testPrintf() {
	s := "testPrintf"
	fmt.Printf("%s", s)
}

func testFprintln() {
	// Fprintln(w io.Writer, a ...interface{}) (n int, err error)
	n, err := fmt.Fprintln(os.Stdout, "testFprintln")
	fmt.Println(n, err)
}

func testFprint() {
	// Fprint(w io.Writer, a ...interface{}) (n int, err error)
	n, err := fmt.Fprint(os.Stdout, "testFprint")
	fmt.Println(n, err)
}

func testFprintf() {
	n, err := fmt.Fprintf(os.Stdout, "%s", "testFprintf")
	fmt.Println(n, err)
}

func testSprintln() {
	// Sprintln(a ...interface{}) string
	s := fmt.Sprintln("testSprintln")
	fmt.Println(s)
}

func testSprint() {
	s := fmt.Sprint("testSprint")
	fmt.Println(s)
}

func testSprintf() {
	s := fmt.Sprintf("%s", "testSprintf")
	fmt.Println(s)
}

func main() {
	n, _ := fmt.Fprint(os.Stdout, "内容") // 向流中写入内容
	fmt.Println(n)
	n2, _ := fmt.Fprintln(os.Stdout, "内容2") // 向流中写入内容,并加上换行符
	fmt.Println(n2)
	n3, _ := fmt.Fprintf(os.Stdout, "%s \n", "内容3") // 格式化向流中写入内容
	fmt.Println(n3)

	fmt.Print("1", "2")           // 输出内容不换行
	fmt.Println("1", "2")         // 输出内容并换行符
	fmt.Printf("%s-%s", "1", "2") // 格式化输出

	ret := fmt.Sprint("1", "2")
	fmt.Println(ret)
	ret2 := fmt.Sprintln("1", "2")
	fmt.Println(ret2)
	ret3 := fmt.Sprintf("%s-%s", "1", "2")
	fmt.Println(ret3)
}
