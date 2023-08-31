package main
import "fmt"
func main() {
	/*//nil
	 var a[]int //声明int类型切片
	 var b=[]int{} //声明并且初始化
	 c :=make([]int,0)
	if a == nil {
		fmt.Println("a是一个nil")
	}
	fmt.Println(a,len(a),cap(a))
	if b== nil{
		fmt.Println("b是一个nil")
	} 
	fmt.Println(b,len(b),cap(b))
	if c==nil{
		fmt.Println("c是一个nil")
	}
		fmt.Println(c,len(c),cap(c))
	*/
	//切片的copy
	/*
	a :=[]int {1,2,3,4,5}
	b:=make([]int,5,5)
	c:=b
	copy(b,a)
	b[0]=100
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	*/
	
	//切片的遍历
	a:=[]int{1,2,3,4,5}
	//使用索引遍历
	for i:=0;i<len(a);i++{
		fmt.Println(i,a[i])
	}	
	fmt.Println()
	//for range 遍历
	for index,value:=range a{
		fmt.Println(index,value)
	}
	
	
}