package utils

import "fmt"

type FamilyAccount struct {
	flag    bool
	check   string //接收退出的指令
	key     string
	balance float64 //初始金额
	money   float64 //每次收支金额
	note    string  //收支说明
	deal    bool    //定义是否有收支行为
	details string  //收支详细记录
}

// 编写要给工厂模式的构造方法，返回一个*FamilyAccount实例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		deal:    false,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收入\t账户余额\t收支金额\t说 明",
	}
}

// 给结构体绑定方法
// 将显示写成一个方法
func (this *FamilyAccount) showDetails() {
	if this.deal == true {
		fmt.Println("----------当前收支明细记录-----------")
		fmt.Println(this.details)
	} else {
		fmt.Println("----------还没有收支记录,快来一笔吧-----------")
	}
}

// 收入方法
func (this *FamilyAccount) income() {
	fmt.Println("本次收入:")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明:")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	this.deal = true
}

// 支出方法
func (this *FamilyAccount) expend() {
	fmt.Println("----------当前支出记录-----------")
	fmt.Println("本次支出:")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足！")
	} else {
		this.balance -= this.money
	}
	fmt.Println("本次支出说明:")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	this.deal = true
}

// 转账方法
func (this *FamilyAccount) transfer() {
	for {
		trans := ""   //接收转入或者转出的指令
		tmoney := 0.0 //接收转账交易金额
		fmt.Println("请选择需要转入还是转出（1/2）:")
		fmt.Scanln(&trans)
		if trans == "1" {
			fmt.Println("请输入需要转入的金额：")
			fmt.Scanln(&tmoney)
			this.note = "入账"
			this.balance += tmoney
			this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", this.balance, tmoney, this.note)
			this.deal = true
			break
		} else if trans == "2" {
			fmt.Println("请输入需要转出的金额：")
			fmt.Scanln(&tmoney)
			this.balance -= tmoney
			this.note = "转出"
			this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", this.balance, tmoney, this.note)
			this.deal = true
			break
		} else {
			fmt.Println("选择错误，请重新选择(1/2)转入还是转出:")
		}
	}
}

// 退出方法
func (this *FamilyAccount) exit() {
	fmt.Println("您真的要退出吗?(输入Y/N表示确否):")
	for {
		fmt.Scanln(&this.check)
		if this.check == "Y" || this.check == "N" {
			break
		}
		fmt.Println("输入有误重新输入:")
	}
	if this.check == "Y" {
		this.flag = true
	} else {
		this.flag = false
	}
}

// 显示主菜单
func (this *FamilyAccount) MainMenu() {
	account := ""
	pw := ""
	times := 3
	for {
		fmt.Println("请输入账号:")
		fmt.Scanln(&account)
		fmt.Println("请输入密码:")
		fmt.Scanln(&pw)
		if account == "136347968" && pw == "ganjing" {
			fmt.Println("登录成功！")
			break
		} else {
			times--
			if times == 0 {
				fmt.Println("验证次数用完，程序即将退出！")
				return
			} else {
				fmt.Printf("账号或密码输入错误，请重新输入！您还有%v次机会\n", times)
			}
		}
	}
	for {
		fmt.Println("\n----------家庭收支软件-----------")
		fmt.Println("----------1.收支明细-----------")
		fmt.Println("----------2.登记收入-----------")
		fmt.Println("----------3.登记支出-----------")
		fmt.Println("----------4.转账功能-----------")
		fmt.Println("----------5.退出程序-----------")
		fmt.Println("输入(1-4):")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.expend()
		case "4":
			this.transfer()
		case "5":
			this.exit()
		}
		if this.flag == true {
			break
		}
	}
	fmt.Println("退出成功！")

}
