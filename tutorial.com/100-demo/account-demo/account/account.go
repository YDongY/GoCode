package account

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

func (a *account) showDetails() {
	fmt.Println("\n------------------当前收支明细-------------")
	if a.flag {
		fmt.Println(a.detail)
	} else {
		fmt.Println("你还没有收支...来一笔吧")
	}
}

func (a *account) income() {

	fmt.Println("本次输入金额：")
	fmt.Scanln(&a.money)
	a.balance += a.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&a.note)
	a.detail += fmt.Sprintf("\n收支\t%v\t%v\t%v", a.balance, a.money, a.note)
	a.flag = true
}

func (a *account) pay() {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&a.money)
	if a.money > a.balance {
		fmt.Println("支出的余额不足")
	} else {
		a.balance -= a.money
		fmt.Println("本次支出说明：")
		fmt.Scanln(&a.note)
		a.detail += fmt.Sprintf("\n支出\t%v\t%v\t%v", a.balance, a.money, a.note)
		a.flag = true
	}

}

func (a *account) exit() {
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
		a.loop = false
	}
}

func (a *account) MainMenu() {
	for {
		fmt.Println("\n-------------------收支软件-----------------")
		fmt.Println("                  1 收支明细                ")
		fmt.Println("                  2 登记收入                ")
		fmt.Println("                  3 登记支出                ")
		fmt.Println("                  4 退出软件                ")
		fmt.Printf("请选择(1-4):")
		fmt.Scanln(&a.key)
		switch a.key {
		case "1":
			a.showDetails()
		case "2":
			a.income()
		case "3":
			a.pay()
		case "4":
			a.exit()
		default:
			fmt.Println("请输入正确选项...")
		}
		if !a.loop {
			break
		}
	}
	fmt.Println("已成功退出软件...")
}