package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	// 将字符串转换为小写，并去除非字母和数字的字符
	lowercase := strings.ToLower(s)
	filtered := ""
	for _, char := range lowercase {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			filtered += string(char)
		}
	}
	// 使用双指针法判断是否回文
	left := 0
	right := len(filtered) - 1
	for left < right {
		if filtered[left] != filtered[right] {
			return false
		}
		left++
		right--
	}
	return true
}
func main() {
	str := "A man, a plan, a canal: Panama"
	if isPalindrome(str) {
		fmt.Println("是回文串")
	} else {
		fmt.Println("不是回文串")
	}
}
