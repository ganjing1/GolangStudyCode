 package main
import (
	"fmt"
"unsafe"
)
func main() {/*
num:= [...]int {2,7,11,15 }
target:= 9
for i:=0;i<len(num);i++{
	for j:=i;j<len(num);j++{
		if num[i]+num[j]==target{
			fmt.Println(i,j)
		}
	}
}*/
//如何在程序上查看某变量占用字节大小和数据类型
var n2 int64 = 10 
//unsafe.Sizeof(n2) 是nsafe包的一个函数,可以返回n2变量占用的字节数
fmt.Printf("n2 的类型 %T n2占用的字节数是 %d\n"	,n2,unsafe.Sizeof(n2))
//整型变量在使用时，遵守保小不保大的原则
//在确保正确运行程序的同时，尽量使用占用字节小的数据类型
var age byte=90
fmt.Println(age)
}