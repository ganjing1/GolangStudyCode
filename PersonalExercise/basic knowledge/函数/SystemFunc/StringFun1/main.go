package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
	str:="hello你好"
	fmt.Println(len(str))

	//利用健值循环;for-range
	for i,value:=range str{
		fmt.Printf("%d,%c",i,value)
	}

	//利用r:=[]rune(str),切片
	r:=[]rune(str)
	for i :=0;i<len(r);i++{
		fmt.Printf("%c\n",r[i])
	}
	//字符串转整数
	num1,_:=strconv.Atoi("666")
	fmt.Println(num1)

	//整数转字符串
	s:=strconv.Itoa(88)
	fmt.Println(s)

	//统计一个字符串中有多少个指定的子串
	i:=strings.Count("gogogogo","go")
	fmt.Println(i)

	//统计子串是否在母串中
	s1:=strings.Contains("gogo","a")
	fmt.Println(s1)

	//不区分字符串大小比较字符串是否相等
	s2:=strings.EqualFold("HELLO","hello")
	fmt.Println(s2)

	//区分字符串大小比较字符串是否相等
	fmt.Println("hello"=="HELLO")

	//返回第一个子串在母串中出现的索引位置
	s3:=strings.Index("hello world","wo")
	fmt.Println(s3)
}
