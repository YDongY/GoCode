package service

import (
	"github.com/project2/model"
)

type CustomerService struct {
	customers []model.Customer
	// 当前切片含有多少客户
	customerNum int
}

func NewCustomerService() *CustomerService {
	cs := &CustomerService{}
	cs.customerNum = 1
	customer := model.NewCustomer(1, "tom", "男", 20, "18888888888", "tom@gmail.com")
	cs.customers = append(cs.customers, customer)

	return cs
}

func (this *CustomerService) List() []model.Customer {
	return this.customers
}

func (this *CustomerService) Add(customer model.Customer) bool {
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

func (this *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = i
			break
		}
	}
	return index
}

func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}
