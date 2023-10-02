package main

import "fmt"

func fibonacciRecursive(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
	}
}

func main() {
	n := 10
	fib := fibonacciRecursive(n)
	fmt.Println("递归方法计算结果：", fib)
}
