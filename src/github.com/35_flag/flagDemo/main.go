package main

import (
	"flag"
	"fmt"
)

func main() {
	var user string
	var pwd string
	var host string
	var port int

	// 接收命令行输入 -u 后面的参数
	// 　　　　　　　值　　　参数　默认值　说明
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "p", "", "密码")
	flag.StringVar(&host, "h", "localhost", "主机名")
	flag.IntVar(&port, "P", 3306, "端口号")

	flag.Parse()
	fmt.Printf("user=%v,pwd=%v,host=%v,port=%v", user, pwd, host, port)

}
