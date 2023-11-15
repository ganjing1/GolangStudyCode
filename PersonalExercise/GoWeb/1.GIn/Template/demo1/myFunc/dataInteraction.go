package myFunc

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Student struct {
	Age  int
	Name string
}

func Hello1(context *gin.Context) {
	age := 19
	arr := []int{33, 66, 99}
	flag := true
	username := "lili"
	stu := Student{Age: 22,
		Name: "甘进"}
	now_time := time.Now()
	//将age和arr放入map:
	map_data := map[string]interface{}{
		"age":      age,
		"arr":      arr,
		"flag":     flag,
		"username": username,
		"stu":      stu,
		"now_time": now_time,
	}
	context.HTML(200, "1.demo1/hello1.html", map_data)
}

func Hello4(context *gin.Context) {

}
func Add(num1 int, num2 int) int {
	return num1 + num2
}
