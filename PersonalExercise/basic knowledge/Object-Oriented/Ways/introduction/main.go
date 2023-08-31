package main
import "fmt"
type Person struct {
	name string
}

//这是方法和函数不一样
/*
Person 结构体有一个test()方法，test()和Person是绑定的
a为结构体Person调用的具体对象，用对象来调用test()
*/
func (a Person) test(){
	a.name="lisi"
	fmt.Println(a.name)
}

type student struct {
	name string
}

func main(){
	var p Person
	//var a student
	p.name="zhangsan"
	p.test()
	//a.test(),错误，因为 只能是结构体Person对应的对象才能调用
	fmt.Println(p.name)
}
