package myfunc

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Hello1(context *gin.Context) {
	context.HTML(200, "1.demo1/hello.html", nil)
}

type Databind struct {
	Username string `form:"username"`
	Pwd      string `form:"pwd"`
}

func Hello2(context *gin.Context) {
	var databind Databind
	err := context.ShouldBind(&databind)
	fmt.Println(databind)
	if err != nil {
		context.String(404, "绑定失败")
	} else {
		context.String(200, "绑定成功")
	}
}
