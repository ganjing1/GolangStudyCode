package main

import (
	"fmt"
)

func findMax(lst []int) int {
	if len(lst) == 1 {
		return lst[0]
	} else {
		mid := len(lst) / 2 //2
		/*
			递归调用 findMax 函数来找到左半部分和右半部分的最大值，
			并将结果分别存储在 leftMax 和 rightMax 变量中
		*/
		leftMax := findMax(lst[:mid])
		rightMax := findMax(lst[mid:])
		if leftMax > rightMax {
			return leftMax
		} else {
			return rightMax
		}
	}
}

func main() {
	list := []int{9, 6, 11, 20, 8}
	fmt.Println(findMax(list))
}
