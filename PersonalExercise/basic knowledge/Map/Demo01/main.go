package main

import "fmt"

func main(){
	var a map[int]string
	//只声明map内存没有分配空间
	a=make(map[int]string,10)//map存放10个键值对,用make创建十个键值对
	a[1001]="ganjing"
	a[1002]="zhangsan"
	a[1003]="Lisi"
	a[1004]="Wangwu"
	a[1001]="ganqizhi"//替换了上面的ganjing
	fmt.Println(a)

	//第二种
	var b map[int]string
	b=make(map[int]string)
	b[01]="hello"
	b[02]="world"
	b[1003]="Lisi"
	b[1004]="Wangwu"
	b[1001]="ganqizhi"//替换了上面的ganjing
	fmt.Println(b)

	//第三种
	c:= map[int]string{
		2001:"zhangsan",
		2002:"lisi",
	}
	fmt.Println(c)
}
