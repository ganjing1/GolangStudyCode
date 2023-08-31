package main

import "fmt"

// 父类
type Animal struct {
	age    int
	Weight float32
}

// 给Animal定义方法：喊叫
func (an *Animal) yell() {
	fmt.Println("you can yell")
}

// <表现> 方法
func (an *Animal) Show() {
	fmt.Printf("animal's age:%v,animal's weight:%v", an.age, an.Weight)
}

type Person struct {
	//继承
	Animal
	age int //子类和父类都有
}

func (an *Person) Show() {
	fmt.Printf("person->animal's age:%v,animal's weight:%v", an.age, an.Weight)
}
func (p *Person) PersonAction() {
	fmt.Println("Person could study")
}
func main() {
	person := &Person{}         //已经创建指针person，但是没有值
	person.Animal.age = 22      //person.Animal.age
	person.Animal.Weight = 65.0 //person.Animal.Weight
	person.PersonAction()
	person.Show() //Animal中得方法
	//person.yell()
}
