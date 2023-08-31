package main

import (
	"fmt"
)

func main() {
	flag := false //管理是否退出
	check := ""   //接收退出的指令
	key := ""
	balance := 1000.0                //初始金额
	money := 0.0                     //每次收支金额
	note := ""                       //收支说明
	deal := false                    //定义是否有收支行为
	details := "收入\t账户余额\t收支金额\t说 明" //收支详细记录
	for {
		fmt.Println("\n----------家庭收支软件-----------")
		fmt.Println("----------1.收支明细-----------")
		fmt.Println("----------2.登记收入-----------")
		fmt.Println("----------3.登记支出-----------")
		fmt.Println("----------4.退出程序-----------")
		fmt.Println("输入(1-4):")
		fmt.Scanln(&key)
		switch key {
		case "1":
			if deal == true {
				fmt.Println("----------当前收支明细记录-----------")
				fmt.Println(details)
			} else {
				fmt.Println("----------还没有收支记录,快来一笔吧-----------")
			}
		case "2":
			fmt.Println("本次收入:")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("本次收入说明:")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", balance, money, note)
			deal = true
		case "3":
			fmt.Println("----------当前支出记录-----------")
			fmt.Println("本次支出:")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额不足！")
			} else {
				balance -= money
			}
			fmt.Println("本次支出说明:")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", balance, money, note)
			deal = true
		case "4":
			fmt.Println("您真的要退出吗?(输入Y/N表示确否):")
			for {
				fmt.Scanln(&check)
				if check == "Y" || check == "N" {
					break
				}
				fmt.Println("输入有误重新输入:")
			}
			if check == "Y" {
				flag = true
			} else {
				flag = false
			}
		default:
			fmt.Println("请输入正确数字：")
		}
		if flag == true {
			break
		}
	}
	fmt.Println("退出成功！")
}
