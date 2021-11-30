package service

import (
	"golab/example.com/100-demo/custom-demo/model"
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

func (c *CustomerService) List() []model.Customer {
	return c.customers
}

func (c *CustomerService) Add(customer model.Customer) bool {
	c.customerNum++
	customer.Id = c.customerNum
	c.customers = append(c.customers, customer)
	return true
}

func (c *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(c.customers); i++ {
		if c.customers[i].Id == id {
			index = i
			break
		}
	}
	return index
}

func (c *CustomerService) Delete(id int) bool {
	index := c.FindById(id)
	if index == -1 {
		return false
	}
	c.customers = append(c.customers[:index], c.customers[index+1:]...)
	return true
}