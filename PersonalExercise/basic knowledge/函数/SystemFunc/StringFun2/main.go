package main

import (
	"fmt"
	"strings"
)

func main(){
	//字符串替换
	//strings.Repalce(s,s1,s2,n)
	/*
		s代表原字符串
		s1为需要替换的字符串
		s2为替换后的字符串
		n 为替换的个数
	*/
	s1:=strings.Replace("I love golang golang","golang","english",1)
	fmt.Println(s1)

	//按照指定的某个字符，位分割标识，将一个字符串拆分为字符数组
	arr:=strings.Split("go-java-python-c++","-")
	fmt.Println(arr)
	fmt.Printf("%T",arr)

	//大小写字符之间的转换
	fmt.Println(strings.ToUpper("hello"))
	fmt.Println(strings.ToLower("WORLD"))

	//去除左右两边空格
	fmt.Println(strings.TrimSpace("    hello   world   "))

	//去除左右两边指定字符
	fmt.Println(strings.Trim("hello","h"))

	////去除右边指定字符
	fmt.Println(strings.TrimRight("hello","o"))

	//判断字符串是否以某字符串开头
	fmt.Println(strings.HasPrefix("hello","he"))

	//判断字符串是否以某字符串结尾
	fmt.Println(strings.HasSuffix("hello","0"))
}
