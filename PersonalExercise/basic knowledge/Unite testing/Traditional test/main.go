package main

import "fmt"

func add(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func main() {
	value := add(10)
	if value == 55 {
		fmt.Printf("add 正确 返回值=%v,期望值=%v\n", value, 55)
	} else {
		fmt.Printf("add 错误 返回值=%v,期望值=%v\n", value, 55)
	}
}
