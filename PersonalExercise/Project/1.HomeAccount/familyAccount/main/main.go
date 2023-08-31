package main

import (
	"familyAccount/utils"
	"fmt"
)

func main() {
	fmt.Println("此程序通过面向对象的方式完成")
	utils.NewFamilyAccount().MainMenu()
}
