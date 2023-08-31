package main

//for 循环
//1.for 循环的基本写法
/*func main(){
	for i:=0;i<10;i++{
		fmt.Println(i)
	}*/
//2.省略初始语句，但是必须保初始语句后面的分号
 /*var i=0
func main(){
	/*for  ;i<10;i++{
		fmt.Println(i)
		} */
//3.省略初始语句和结束语句
/* var i=10
for i>0{
	fmt.Println(i)
	i-- 
} */
//4.break跳出for循环	
func main(){
	for i:=0;i<5;i++{
		if i==3{
			continue
		}
	}
}

