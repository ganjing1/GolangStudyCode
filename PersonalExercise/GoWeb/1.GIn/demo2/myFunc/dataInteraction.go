package myFunc

import "github.com/gin-gonic/gin"

func Hello(context *gin.Context) {
	name := "my name is Ganjin"
	//200代表状态码：请求成功。一般用于GET与POST请求
	context.HTML(200, "demo1/hello1.html", name)
}

type Student struct {
	Name string
	Age  int
}

func Hello2(context *gin.Context) {
	info := Student{Name: "张三",
		Age: 12}
	//200代表状态码：请求成功。一般用于GET与POST请求
	context.HTML(200, "demo2/hello2.html", info)
}

func Array(context *gin.Context) {
	arr := []int{10, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//200代表状态码：请求成功。一般用于GET与POST请求
	context.HTML(200, "demo3/ArrayDemo.html", arr)
}
func StructArray(context *gin.Context) {
	arr := []Student{{"zhangsan", 18},
		{"ddddgsan", 20},
		{"lili", 21}}
	//200代表状态码：请求成功。一般用于GET与POST请求
	context.HTML(200, "demo4/StructArray.html", arr)
}
func Map(context *gin.Context) {
	Mp := make(map[string]int, 3)
	Mp["zhangsan"], Mp["lisi"], Mp["aaa"] = 101, 202, 303
	context.HTML(200, "demo5/Map.html", Mp)
}
func MapStruct(context *gin.Context) {
	Mp := make(map[int]Student, 3)
	Mp[1001] = Student{"张三", 12}
	Mp[1002] = Student{"lisi", 13}
	Mp[1003] = Student{"ww", 15}
	context.HTML(200, "demo6/MapStruct.html", Mp)
}
func Slice(context *gin.Context) {
	sl := []int{1, 2, 3, 4, 5, 6}
	context.HTML(200, "demo7/Slice.html", sl)
}
