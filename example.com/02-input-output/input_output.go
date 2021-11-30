package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Error() {
	const name, age = "mark", 24
	err := fmt.Errorf("error name=%s,age=%d", name, age)
	fmt.Println(err.Error())
}

func Input() {
	var (
		name string
		age  int
	)
	r := strings.NewReader("mark 24")
	// -------------- Fscan / Fscanf / Fscanln --------------
	n, err := fmt.Fscan(r, &name, &age)
	fmt.Println(n, name, age, err)

	n, err = fmt.Fscanf(r, "%s %d", &name, &age)
	fmt.Println(n, name, age, err)

	// 换行符处停止扫描，并且在最后一项之后必须有一个换行符或 EOF
	n, err = fmt.Fscanln(r, &name, &age)
	fmt.Println(n, name, age, err)

}

func Output() {
	/**
	- Xprint:   直接输出
	- Xprintln: 输出并添加换行
	- Xprintf:  格式化输出

		1. Fxx: 向流中输出数据
		2. Sxx: 向字符串变量输出数据
		3. Pxx: 向 os.Stdout 输出数据
	*/
	// ------------------- Fprint / Fprintf / Fprintln ------------------
	const name, age = "mark", 24
	n, err := fmt.Fprint(os.Stdout, name, " is ", age, " years old.\n")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprint: %v\n", err)
	}
	fmt.Print(n, " bytes written.\n")

	// ================

	n, err = fmt.Fprintf(os.Stdout, "%s is %d years old.\n", name, age)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
	}
	fmt.Printf("%d bytes written.\n", n)

	// ================

	n, err = fmt.Fprintln(os.Stdout, name, "is", age, "years old.")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintln: %v\n", err)
	}
	fmt.Println(n, "bytes written.")

	// ------------------ Print / Printf / Println ----------------------
	fmt.Print(name, " is ", age, " years old.\n")
	fmt.Printf("%s is %d years old.\n", name, age)
	fmt.Println(name, "is", age, "years old.")

	// ----------------- Sprint / Sprintf / Sprintln ----------------------
	s := fmt.Sprint(name, " is ", age, " years old.\n")
	io.WriteString(os.Stdout, s)
	s = fmt.Sprintf("%s is %d years old.\n", name, age)
	io.WriteString(os.Stdout, s)
	s = fmt.Sprintln(name, "is", age, "years old.")
	io.WriteString(os.Stdout, s)
}

func main() {
	Output()
	Input()
	Error()

}
