package main

import (
	"fmt"
	"github.com/07_网络编程/TCPProject/client/process"
	"os"
)

var (
	userId   int
	userPwd  string
	userName string
)

func main() {
	var key int
	for {
		fmt.Println("------------- Menu -------------")
		fmt.Println("\t 1 Login")
		fmt.Println("\t 2 Register")
		fmt.Println("\t 3 Exit")
		fmt.Println("\t Please Select[1-3]")
		fmt.Println("--------------------------------")

		_, _ = fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Printf("【Login ChatRoom】 \n")
			// 登录界面
			fmt.Printf("请输入用户 ID: ")
			_, _ = fmt.Scanf("%d\n", &userId)
			fmt.Printf("请输入用户 Password: ")
			_, _ = fmt.Scanf("%s\n", &userPwd)
			u := &process.UserProcess{}
			u.Login(userId, userPwd)
		case 2:
			fmt.Printf("【Register User】 \n")
			fmt.Printf("请输入用户 ID: ")
			_, _ = fmt.Scanf("%d\n", &userId)
			fmt.Printf("请输入用户 Password: ")
			_, _ = fmt.Scanf("%s\n", &userPwd)
			fmt.Printf("请输入用户 Name: ")
			_, _ = fmt.Scanf("%s\n", &userName)
			u := &process.UserProcess{}
			u.Register(userId, userPwd, userName)
		case 3:
			fmt.Printf("【exit】 \n")
			os.Exit(0)
		default:
			fmt.Printf("【Select error】 \n")
		}
	}
}
