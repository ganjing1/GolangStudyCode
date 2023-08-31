package  main
import"fmt"
import"math"
func main(){
	//十进制打印二进制
	n:=10
	fmt.Println("%b\n",n)	
	fmt.Println("%d\n",n)
	//八进制
	m:=075
	fmt.Println("%d\n",m)
	fmt.Println("%o\n",m)	
	//十六进制
	f:=0xff
	fmt.Println("%x\n",f)
	//uint8
	var age uint8
	age=255
	fmt.Println(age)
	//浮点型
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
	//布尔值
	var a bool
	fmt.Println(a)
	a=true
	fmt.Println(a)
	//字符串
	s1:="hello beijing"
	s2:="你好 北京"
	fmt.Println(s1,s2)
	//多行字符串
	/*s3:= '
	多行字符串
	两个反引号之间的内容
	会原样输出
	\n
	\t
	'
	fmt.Println(s3)*/
}