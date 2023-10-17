package myfunc

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type User2 struct {
	Uname string `json:"uname" uri:"uname" form:"uname"`
	Age   int    `json:"age" uri:"age" form:"age"`
	/*
		后面又有`json，又有uri，又有form是指定不同的标签键和值，
		用来控制不同的包或库如何处理结构体字段。例如，json标签是用来指定encoding/json包如何编码和解码结构体字段的，
		uri标签和form标签可能是用来指定某些Web框架如何从URL或表单中获取结构体字段的。
		不同的标签键和值之间用空格分隔，具体的格式和含义取决于具体的包或库。
	*/
}

func Hello6(context *gin.Context) {
	//创建结构体示例：
	var user User2
	//数据绑定：
	err := context.ShouldBindUri(&user)
	//打印结构体对象的内容：
	fmt.Println(user)
	if err != nil {
		context.String(404, "绑定失败")
	} else {
		context.String(200, "绑定成功")
	}
}
