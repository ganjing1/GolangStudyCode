package main

import (
	"Encapsulation/model"
	"fmt"
)

func main() {
	p := model.NewPerson("zhangsan")
	p.SetAge(20)
	fmt.Println(p.Name)
	fmt.Println(p.GetAge())
	fmt.Println(p)
}
