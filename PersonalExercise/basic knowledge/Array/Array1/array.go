package main

import (
	"fmt"
)

// 数组相关内容
func main() { /*
		var a[3] int
		var b[4] int
		fmt.Println(a)
		fmt.Println(b)
		//数组的初始化
		//1.定义时使用初始值列表的方式初始化
		var cityarray=[4]string{"北京","上海","广州","深圳"}
		fmt.Println(cityarray)
		//2.编译器推导数组的长度
		var boolarray=[...]bool{true,false,false,true,false}
		fmt.Println(boolarray)
		fmt.Println(cityarray[0])
		fmt.Println(cityarray[3])
		//3.使用索引值方式初始化
		var langarray= [...] string{1:"golang",3:"python",7:"java"}
		fmt.Println(langarray[1])
		fmt.Println("%T\n",langarray)*/
	//数组的遍历
	//1.for 循环遍历
	//var cityarray=[4]string{"北京","上海","广州","深圳"}
	/*for i:=0;i<len(cityarray);i++{
		fmt.Println(cityarray[i])
	}*/
	//2.for range遍历
	/*for _,value := range cityarray {
	fmt.Println(value)
	}*/
	//2.二维数组
	/*cityarray:=[4][2]string{
		{"北京","西安"},
		{"武汉","长沙"},
		{"杭州","厦门"},
		{"广州","福州"},
	}
	fmt.Println(cityarray)
	fmt.Println(cityarray[2][0])
	//二维数组的遍历
	for _,v1:=range cityarray{
			for _,v2:=range v1{
				fmt.Println(v2)
			}
		}
	}*/
	//数组是值类型
	x := [3][2]int{{1, 2}, {2, 3}, {3, 4}}
	fmt.Println(x)
	f1(x)
	fmt.Println(x)
}
func f1(a [3][2]int) {
	a[0][0] = 100
}
