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

	fmt.Println(uname)
	fmt.Println(pwd)

}
func Hello3(context *gin.Context) {
	uname := context.PostForm("uname")
	fmt.Println(uname)
	if uname == "ganjing" {
		//方式1
		//mapdata := map[string]interface{}{
		//	"msg": "用户名重复",
		//}
		//context.JSON(200, mapdata)

		//方式2
		context.JSON(200, gin.H{
			"msg": "用户名重复",
		})
	} else {
		context.JSON(200, gin.H{
			"msg": "用户名可用",
		})
	}
}
