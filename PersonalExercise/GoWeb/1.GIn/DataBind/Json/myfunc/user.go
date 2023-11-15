package myfunc

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Hello(context *gin.Context) {
	context.HTML(200, "1.demo1/hello.html", nil)
}

type JsonBind struct {
	Uname string `json:"uname"`
	Age   int    `json:"age"`
}

func Hello2(context *gin.Context) {
	var JB JsonBind
	err := context.ShouldBind(&JB)
	fmt.Println(JB)
	if err != nil {
		context.JSON(404, gin.H{
			"msg": "绑定失败",
		})
	} else {
		context.JSON(200, gin.H{
			"msg": "绑定成功",
		})
	}
}
