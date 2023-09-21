package main

import "fmt"

func Bubble_sort(lst []int) {
	n := len(lst)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if lst[j] > lst[j+1] {
				lst[j], lst[j+1] = lst[j+1], lst[j] //简单互换
			}
		}
	}
	fmt.Println(lst)

}

func main() {
	list := []int{22, 11, 33, 66, 99, 10}
	Bubble_sort(list)

}
