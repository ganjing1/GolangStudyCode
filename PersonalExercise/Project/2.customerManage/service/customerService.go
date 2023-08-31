package service

import "2.customerManage/model"

// 完成对customer的操作，包含增删改查
type CustomerService struct {
	customers   []model.Customer //将每个客户的信息都存放到切片中
	customerNum int
}

// 此方法，返回*customerService
func NewCustomerService() *CustomerService {
	//为了能够看到有客户在切片中，我们初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "zhangsan", "man", 20, "111", "zs@sohu.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

// 返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}
