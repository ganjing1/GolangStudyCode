package myfunc

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type URLbind struct {
	Username string `form:"username"`
	PWD      string `form:"pwd"`
}

func Hello(context *gin.Context) {
	var url URLbind
	err := context.ShouldBind(&url)
	fmt.Println(url)
	if err != nil {
		context.String(400, "绑定失败")
	} else {
		context.String(200, "绑定成功")
	}
}
