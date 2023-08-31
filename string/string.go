package main
import (
	"fmt"	
	"strings"
)
func main() {
	s := "hello"
	fmt.Println(len(s))
	s1:="你好甘进"
	fmt.Println(len(s1))
	//拼接字符串求字符串的长度
	fmt.Println(s+s1)
	s3:=fmt.Sprintf("%s_%s",s,s1)
	fmt.Println(s3)
	//字符串的分割
	s4:="how do you do"
	fmt.Println(strings.Split(s4," "))
	fmt.Println("%T\n",strings.Split(s4," "))
	//判断是否包含
	fmt.Println(strings.Contains(s4,"do"))
	//判断前缀
	fmt.Println(strings.HasPrefix(s4,"how"))
	//判断后缀
	fmt.Println(strings.HasSuffix(s4,"do"))
	//判断子串的位置
	fmt.Println(strings.Index(s4,"do"))
	//最后字符串出现的位置
	fmt.Println(strings.LastIndex(s4,"do"))
	//join 操作
	s5:=[]string{"how","do","you","do"}
	fmt.Println(s5)
	fmt.Println(strings.Join(s5,"-"))


}