package main

import (
	"fmt"
)

func Binary_search(lst []int, target int) {
	n := len(lst)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if lst[j] > lst[j+1] {
				lst[j], lst[j+1] = lst[j+1], lst[j]
			}
		}
	}
	fmt.Printf("排序后的数组为:%v\n", lst)
	var mid int
	flag := 0
	high := n - 1
	low := 0
	for high >= low {
		mid = (low + high) / 2

		if lst[mid] > target {
			high = mid - 1

		} else if lst[mid] < target {
			low = mid + 1
		} else {
			flag = 1
			break
		}
	}
	if flag == 1 {
		fmt.Printf("%v在数组中的%v号位置上\n", target, mid)
	} else {
		fmt.Println("未找到", target)
	}
}
func main() {
	list := []int{22, 55, 66, 4, 12, 33, 66, 11}
	Binary_search(list, 55)
}
