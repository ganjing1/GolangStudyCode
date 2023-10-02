package main

import (
	"fmt"
	"strings"
)

func countWords(s string) int {
	// 使用空格分割字符串
	words := strings.Split(s, " ")

	// 统计单词个数
	count := len(words)

	return count
}

func main() {
	s := "Hello World"

	count := countWords(s)

	fmt.Println("单词个数：", count)
}
