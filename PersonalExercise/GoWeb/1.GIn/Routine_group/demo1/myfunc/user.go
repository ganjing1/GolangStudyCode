package myfunc

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `Form:"username"`
	Password string `Form:"pwd"`
}

func Hello1(context *gin.Context) {
	context.HTML(200, "1.demo1/hello1.html", nil)
}
func Hello2(context *gin.Context) {
	var databind User
	err := context.ShouldBind(&databind)
	fmt.Println(databind)
	if err != nil {
		context.String(404, "绑定失败")
	} else {
		context.String(200, "绑定成功")
	}
}
func Hello3(context *gin.Context) {}

func Hello4(context *gin.Context) {}
func Hello5(context *gin.Context) {}
func Hello6(context *gin.Context) {}
