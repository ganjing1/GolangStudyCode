package main

import (
	"fmt"	
	"sort"
)

func main() {
var a=make([]string,5,10)
for i:=1;i<10;i++{
	a=append(a, fmt.Sprintf("%v",i))
}
fmt.Println(a)
	var b=[...]int{1,2,3,4,56,7,89}
	sort.Ints(b[:])//b[:]得到的是一个切片，指向了底层的数组a
	fmt.Println(b)
}