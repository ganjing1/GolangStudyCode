package main

import "fmt"

func main(){
	//定义map
	b:= map[int]string{
		1001:"hello",
		1002:"world",
		1003:"golang",

	}
	b[1001]="goodbye"//修改
	delete(b,1003)//删除
	fmt.Println(b)
	c1,c2:=b[1001]//','号前面表示键， ‘,’号后面表示值
	fmt.Println(c1)//键
	fmt.Println(c2)//值

	fmt.Println(len(b))

	//遍历map
	for k,v:=range b{
		fmt.Printf("%v,%v\n",k,v)
	}

	//复杂map
	//定义一个map a,a的键为map[string],a的值为map[int]string
	a:=make(map[string]map[int]string)
	a["class1"]=make(map[int]string,3)//a中的每个键，对应三个键值
	a["class1"][1001]="aa"
	a["class1"][1002]="bb"
	a["class1"][1003]="cc"

	a["class2"]=make(map[int]string,3)
	a["class2"][2001]="dd"
	a["class2"][2002]="ee"
	a["class2"][2003]="ff"

	for v1,v2:=range a{
		fmt.Println(v1)
		for k1,k2:=range v2{
			fmt.Printf("%v，%v\t",k1,k2)
		}
		fmt.Println()
	}
	//外层循环以a作为遍历对象，内存循环以a对应的value作为遍历对象
}
