package main

import (
	"fmt"
)

func findMaxElement(nums []int) int {
	max := 0
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
func main() {
	numbers := []int{1, 5, 3, 9, 2}
	maxNumber := findMaxElement(numbers)
	fmt.Println(maxNumber) // 输出：9
}
