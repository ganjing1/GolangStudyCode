package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	// 使用空格分割字符串
	words := strings.Split(s, " ")

	// 反转单词列表
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// 拼接反转后的单词列表
	result := strings.Join(words, " ")

	return result
}

func main() {
	s := "Hello World"

	result := reverseWords(s)

	fmt.Println("反转后的字符串：", result)
}
