package main

import (
	model2 "Encapsulation/model"
	"fmt"
)

func main() {
	model := new(model2.Model)
	model.Mymodule()
	var name string
	fmt.Scan(&name)
}
