package main

import (
	"2.customerManage/service"
	"fmt"
)

type customerView struct {
	key  string //接收customer的输入
	loop bool   //表示循环的显示菜单
	//增加一个字段customerService
	customerView *service.CustomerService
}

// 显示所有的客户信息
func (this *customerView) list() {
	//首先，获取到当前所有的客户信息(在切片中)
	customers := this.customerService.List()
	//显示
	fmt.Println("-------------客户列表--------------")
	fmt.Println("编号\t姓名	性别		年龄		电话		邮箱")
	for i := 0; i < len(customers); i++ {

	}

	fmt.Println("-------------客户列表完成--------------")
}

// 显示主菜单
func (this *customerView) mainMenu() {
	for {
		fmt.Println("----------Customer's Info Manage Software---------")
		fmt.Println("----------1.Add customer---------")
		fmt.Println("----------2.Revise customer---------")
		fmt.Println("----------3.Delete customer---------")
		fmt.Println("----------4.Customer list-----------")
		fmt.Println("----------5.Exit-----------")
		fmt.Print("please choice(1-5):")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			fmt.Println("1.Add customer")
		case "2":
			fmt.Println("2.Revise customer")
		case "3":
			fmt.Println("3.Delete customer")
		case "4":
			fmt.Println("4.Customer list")
		case "5":
			this.loop = false
		default:
			fmt.Println("your input have error ,please input again...")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("you have exited the customer manage system!")
}

func main() {
	customerView := customerView{
		key:  "",
		loop: true,
	}
	//这里完成对customerView结构体的customerService字段的初始化
	customerView.customerService = service.NewCustomerService()
	//显示主菜单
	customerView.mainMenu()

}
