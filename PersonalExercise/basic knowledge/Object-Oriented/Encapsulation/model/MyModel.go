package model

import "fmt"

type Model struct {
}

func (m *Model) Mymodule() {
	fmt.Printf("我是Mymodel方法")
}

// import (
//
//	"fmt"
//
// )
type person struct {
	Name string
	age  int //其他包不能访问
}

// 定义工厂模式的函数
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

// 定义set和get方法，对age字段进行封装，在方法中可以加一系列的限制操作，
// 确保被封装字段的安全合理性
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("传入年龄不正确")
	}
}

func (p *person) GetAge() int {
	return p.age
}
