package main

import "fmt"

func main(){
	type Myint int//将int起了一个别名called MyInt
	var a Myint=22
	fmt.Println(a)

	//var b int=30
	//b=a
	//虽然是别名。但是在go中编译识别的时候还是认为
	// Myint和int不一样
	fmt.Println(test(10,20))

}

func test(num1 int ,num2 int)(sub int,add int){
	sub=num1-num2
	add=num1+num2
	return
}
