package myFunc

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Hello1(context *gin.Context) {
	context.HTML(200, "demo1/hello1.html", nil)
}
func Hello2(context *gin.Context) {

	uname := context.PostForm("username")
	pwd := context.PostForm("pwd")
	age := context.DefaultPostForm("age", "18")
	hobby := context.PostFormArray("hobby")
	usermap := context.PostFormMap("user")
	fmt.Println(uname)
	fmt.Println(pwd)
	fmt.Println(age)
	fmt.Println(hobby)
	fmt.Println(usermap)
}
