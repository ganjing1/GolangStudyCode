package myFunc

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type User2 struct {
	Uname string `json:"uname"`
	Age   int    `json:"age"`
}

func Hello4(context *gin.Context) {

	context.HTML(200, "demo1/hello3.html", nil)
}
func Hello5(context *gin.Context) {
	var user User2
	//数据绑定
	err := context.ShouldBind(&user)
	fmt.Println(user)
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
