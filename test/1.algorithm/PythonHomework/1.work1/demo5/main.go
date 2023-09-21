package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := 0
	var i int
	i = 0
	for _, i = range arr {
		sum += i
	}
	fmt.Printf("%v", sum)

}
