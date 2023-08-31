package main

import "fmt"

func main() {
	for {
		choose := -1
		fmt.Println("——————客户信息管理系统——————")
		fmt.Println("1 添加客户       2 修改客户")
		fmt.Println("3 删除客户       4 客户列表")
		fmt.Println("5 退出")
		fmt.Println("请选择 (1~5)")
		fmt.Scanln(&choose)
		if choose != 5 {
			if choose == 1 {
				addCus()
			} else if choose == 2 {
				updateCus()
			} else if choose == 3 {
				delCus()
			} else if choose == 4 {
				listCus()
			} else {
				return
			}
		} else {
			return
		}
	}

}

// 定义客户结构体
type customer struct {
	id    string
	name  string
	sex   string
	age   int
	phone string
	email string
}

// 存储客户信息
var cusList []customer

// 添加
func addCus() {
	var addCus customer
	fmt.Println("————添加客户————")
	fmt.Print("编号： ")
	fmt.Scanln(&addCus.id)
	fmt.Print("姓名： ")
	fmt.Scanln(&addCus.name)
	fmt.Print("性别： ")
	fmt.Scanln(&addCus.sex)
	fmt.Print("年龄： ")
	fmt.Scanln(&addCus.age)
	fmt.Print("电话： ")
	fmt.Scanln(&addCus.phone)
	fmt.Print("邮箱： ")
	fmt.Scanln(&addCus.email)
	fmt.Println("————添加完成————")
	cusList = append(cusList, addCus)
}

// 修改
func updateCus() {
	var updateCus customer
	var upID string
	fmt.Println("————修改客户————")
	fmt.Print("请输入待修改的客户编号(-1退出到菜单)： ")
	fmt.Scanln(&upID)
	if upID != "-1" {
		for index, value := range cusList {
			if value.id == upID {
				fmt.Println("————————修改客户信息(直接回车不修改)———————")
				fmt.Println("编号：", value.id)

				fmt.Printf("姓名(%s)：", value.name)
				fmt.Scanln(&updateCus.name)
				if updateCus.name != "" {
					cusList[index].name = updateCus.name
				}

				fmt.Printf("性别(%s)：", value.sex)
				fmt.Scanln(&updateCus.sex)
				if updateCus.sex != "" {
					cusList[index].sex = updateCus.sex
				}

				fmt.Printf("年龄(%d)：", value.age)
				fmt.Scanln(&updateCus.age)
				if updateCus.age != 0 {
					cusList[index].age = updateCus.age
				}

				fmt.Printf("电话(%s)：", value.phone)
				fmt.Scanln(&updateCus.phone)
				if updateCus.phone != "" {
					cusList[index].phone = updateCus.phone
				}

				fmt.Printf("邮箱(%s)：", value.email)
				fmt.Scanln(&updateCus.email)
				if updateCus.email != "" {
					cusList[index].email = updateCus.email
				}
				fmt.Println("—————————修改完成————————————")
				break
			}
		}
		fmt.Println("输入编号不存在")
	} else {
		return
	}
}

// 删除
func delCus() {
	var delID string
	var choose string
	fmt.Println("————删除客户————")
	fmt.Print("请输入待删除的客户编号(-1退出)： ")
	fmt.Scanln(&delID)
	if delID != "-1" {
		for index, value := range cusList {
			//存在输入的id
			if value.id == delID {
				fmt.Println("编号：", value.id)
				fmt.Println("姓名：", value.name)
				fmt.Println("性别：", value.sex)
				fmt.Println("年龄：", value.age)
				fmt.Println("电话：", value.phone)
				fmt.Println("邮箱：", value.email)
				fmt.Println("确认删除：(是)Y   (否)N")
				fmt.Scanln(&choose)
				if choose == "Y" || choose == "y" {
					//删除的元素在结尾
					if index == len(cusList)-1 {
						cusList = cusList[:index] //将原来的切片大小变成只取到cusList[0:index-1]
						return
						//删除的为首元素
					} else if index == 0 {
						cusList = cusList[1:] //列表的第一个元素删除。这里使用了切片的截取操作，将列表从第二个元素开始截取到末尾，即删除了第一个元素
						return
					} else {
						//删除的元素在中间：将元素的前面和后面重新拼接，不包括要删除的元素本身
						cusList = append(cusList[:index], cusList[index+1:]...) //如果没有...,那么输出的时候就变成两个切片了，有...就表示将cusList[index+1:]逐个元素放进cusList中
						return
					}
				} else {
					return
				}
			}
		}
		// 输入的id没有
		fmt.Println("编号输入有误！")
		return
	} else {
		return
	}

}

// 查询
func listCus() {
	var findID string
	fmt.Println("输入 id 查找，回车显示全部")
	fmt.Scanln(&findID)
	if len(cusList) == 0 {
		fmt.Println("暂无客户信息，返回菜单")
		return
	} else {
		//全部查询
		if findID == "" {
			fmt.Println("————————客户列表——————————")
			fmt.Println("编号", "\t姓名", "\t性别", "\t年龄", "\t电话", "\t邮箱")
			for _, value := range cusList {
				fmt.Println(value.id, "\t", value.name, "\t", value.sex, "\t", value.age, "\t", value.phone, "\t", value.phone)
			}
			fmt.Println("——————————————————————")
			//根据ID查询
		} else {
			for _, value := range cusList {
				if value.id == findID {
					fmt.Println("————————客户信息———————")
					fmt.Println("编号： ", value.id)
					fmt.Println("姓名： ", value.name)
					fmt.Println("性别： ", value.sex)
					fmt.Println("年龄： ", value.age)
					fmt.Println("电话： ", value.phone)
					fmt.Println("邮箱： ", value.email)
					fmt.Println("——————————————————————")
					return
				}
			}
			fmt.Println("输入ID不存在，返回菜单")
			return
		}
	}
}
