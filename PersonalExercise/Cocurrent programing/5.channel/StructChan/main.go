package main

import "fmt"

type Cat struct {
	Name   string
	Age    int
	Weight float64
}

func main() {
	var catChan chan Cat
	catChan = make(chan Cat, 10)
	cat1 := Cat{Name: "kitty", Age: 10, Weight: 11.1}
	cat2 := Cat{Name: "Hitty", Age: 20, Weight: 12.1}
	catChan <- cat1
	catChan <- cat2
	cat3 := <-catChan
	cat4 := <-catChan
	fmt.Println(cat3, cat4)
}
