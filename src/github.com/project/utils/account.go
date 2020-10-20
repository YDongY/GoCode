package utils

import "fmt"

type account struct {
	key     string
	loop    bool
	balance float64
	money   float64
	note    string
	flag    bool
	detail  string
}

func NewAccount() *account {
	return &account{
		key:     "",
		loop:    true,
		balance: 10000.0,
		note:    "",
		flag:    false,
		detail:  "收支\t账户余额\t收支金额\t说明",
	}
}

func (this *account) showDetails() {
	fmt.Println("\n------------------当前收支明细-------------")
	if this.flag {
		fmt.Println(this.detail)
	} else {
		fmt.Println("你还没有收支...来一笔吧")
	}
}

func (this *account) income() {

	fmt.Println("本次输入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	this.detail += fmt.Sprintf("\n收支\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *account) pay() {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("支出的余额不足")
	} else {
		this.balance -= this.money
		fmt.Println("本次支出说明：")
		fmt.Scanln(&this.note)
		this.detail += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)
		this.flag = true
	}

}

func (this *account) exit() {
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
		this.loop = false
	}
}

func (this *account) MainMenu() {
	for {
		fmt.Println("\n-------------------收支软件-----------------")
		fmt.Println("                  1 收支明细                ")
		fmt.Println("                  2 登记收入                ")
		fmt.Println("                  3 登记支出                ")
		fmt.Println("                  4 退出软件                ")
		fmt.Printf("请选择(1-4):")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确选项...")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("已成功退出软件...")
}
