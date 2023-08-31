package main
import( "fmt"
"strconv"
)
func main(){
	//在程序开发中我们经常将基本数据类型转成string，或者将string转换成基本数据类型。
   //基本数据类型和string的转换
	   var num1 int =99
	   var num2 float64=23.456
	   var   b bool=true
	 var mychar byte='h'
	   var str string//空的str
	   //使用第一种方式来转换fmt.sprintf方法
	   str = fmt.Sprintf("%d",num1)
	   fmt.Printf("str type %T str=%q\n",str,str)

	   str = fmt.Sprintf("%f",num2)
	   fmt.Printf("str type %T str=%q\n",str,str)

	   str = fmt.Sprintf("%t",b)//%t代表是按照bool值输出
	   fmt.Printf("str type %T str=%q\n",str,str)//%q该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示

	   str = fmt.Sprintf("%c",mychar)
	   fmt.Printf("str type %T str=%q\n",str,str)


	   //第二种方式 strconv 函数
	   var num3 int =99
	   var num4 float64=23.456
	  var   b2 bool=true
	   

	   str=strconv.FormatInt(int64(num3),10)
	   fmt.Printf("str type %T str=%q\n",str,str)
	   //str=strconv.FormatInt(num4,'f',10,64)
	   //说明：‘f’ 格式 10：表示小数位保留10位 64 ：表示这个小数是float64 
	   str=strconv.FormatFloat(num4,'f',10, 64)
	   fmt.Printf("str type %T str=%q\n",str,str)

	   str=strconv.FormatBool(b2) 
	   fmt.Printf("str type %T str=%q\n",str,str)

	   var num10 int64 = 4567
	   str = strconv.Itoa(int(num10))
	   fmt.Printf("str type %T str=%q\n",str,str)
   
   }

