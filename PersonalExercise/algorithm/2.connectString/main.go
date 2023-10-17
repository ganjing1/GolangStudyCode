package main

import "fmt"

func join(s []string) string {
	var St string
	for _, v := range s {
		St += v
	}
	return St
}

func main() {
	a := []string{
		"hello", "+", "world"}
	fmt.Println(join(a))
}
