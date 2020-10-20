package main

import (
	"fmt"

	"github.com/project2/model"
	"github.com/project2/service"
)

type customerView struct {
	key             string
	loop            bool
	customerService *service.CustomerService
}

func (this *customerView) list() {
	customers := this.customerService.List()
	fmt.Println("\n-------------------客户列表---------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for _, c := range customers {
		fmt.Println(c.GetInfo()) // 获取当前客户信息
	}
}

func (this *customerView) add() {
	fmt.Println("\n-------------------添加客户---------------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱:")
	email := ""
	fmt.Scanln(&email)

	customer := model.NewCustomer2(name, gender, age, phone, email)
	this.customerService.Add(customer)
}

func (this *customerView) delete() {
	fmt.Println("----------------删除客户---------------")
	fmt.Println("请选择删除客户编号(-1退出):")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}

	fmt.Println("你确定要删除吗？y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "Y" || choice == "y" {
			if this.customerService.Delete(id) {
				fmt.Println("----------删除完成------------")
			} else {
				fmt.Println("-----------输入的Id不存在--------")
			}
			break
		} else if choice == "N" || choice == "n" {
			break
		}
		fmt.Println("你输得有误，请重新输入 y/n")
	}

}

func (this *customerView) exit() {
	fmt.Println("你确定要删除吗？y/n")
	for {
		fmt.Scanln(&this.key)
		if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n" {
			break
		}
		fmt.Println("请重新输入(y/n)")
	}

	if this.key == "Y" || this.key == "y" {
		this.loop = false
	}
}

func (this *customerView) mainMenu() {
	for this.loop {
		fmt.Println("\n------------------客户信息---------------")
		fmt.Println("                  1 添加客户                ")
		fmt.Println("                  2 修改客户                ")
		fmt.Println("                  3 删除客户                ")
		fmt.Println("                  4 客户列表                ")
		fmt.Println("                  5 退   出                ")
		fmt.Printf("请选择(1-5):")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			fmt.Println("修改客户")
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("输入有误")
		}
	}
	fmt.Println("已成功退出...")

}

func main() {
	cv := customerView{
		key:  "",
		loop: true,
	}
	cv.customerService = service.NewCustomerService()
	cv.mainMenu()
}
