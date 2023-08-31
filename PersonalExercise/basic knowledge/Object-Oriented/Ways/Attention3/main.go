package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s *Student) String() string { /*
		定义了一个名叫Stirng的方法，该方法可以被外包访问，
		*Student是一个指向Student类型的指针，
		而s是一个指向Student类型的变量
		最后的stirng为返回类型
	*/
	str := fmt.Sprintf("Name=%v,age=%v", s.Name, s.Age)
	return str

}

func main() {
	stu := Student{ /*
			让结构体指针student所指向的 Name为aa，Age为20，最后把student赋值给stu
			stu也是结构体指针
		*/
		Name: "aa",
		Age:  20,
	}
	fmt.Printf("%T", stu)
	fmt.Println()
	fmt.Println((&stu))

}
