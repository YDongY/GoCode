package main

import (
	"fmt"
	"golab/example.com/100-demo/account-demo/account"
)

func main() {
	fmt.Println("面向对象")
	account.NewAccount().MainMenu()
}

func main2() {
	key := ""
	loop := true

	// 账户余额
	balance := 10000.0
	money := 0.0
	note := ""
	flag := false
	detail := "收支\t账户余额\t收支金额\t说明"
	for {
		fmt.Println("\n-------------------收支软件-----------------")
		fmt.Println("                  1 收支明细                ")
		fmt.Println("                  2 登记收入                ")
		fmt.Println("                  3 登记支出                ")
		fmt.Println("                  4 退出软件                ")
		fmt.Printf("请选择(1-4):")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("\n------------------当前收支明细-------------")
			if flag {
				fmt.Println(detail)
			} else {
				fmt.Println("你还没有收支...来一笔吧")
			}

		case "2":
			fmt.Println("本次输入金额：")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			detail += fmt.Sprintf("\n收支\t%v\t%v\t%v", balance, money, note)
			flag = true
		case "3":
			fmt.Println("本次支出金额：")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("支出的余额不足")
				break
			}
			balance -= money
			fmt.Println("本次支出说明：")
			fmt.Scanln(&note)
			detail += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
			flag = true
		case "4":
			fmt.Println("你确定要退出吗？y/n")
			choice := ""
			for {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n" {
					break
				}
				fmt.Println("你输得有误，请重新输入 y/n")
			}
			if choice == "y" {
				loop = false
			}
		default:
			fmt.Println("请输入正确选项...")
		}
		if !loop {
			break
		}
	}

	fmt.Println("已成功退出软件...")

}