package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func _query(c *gin.Context) {

	fmt.Println(c.GetQuery("user"))
	fmt.Println(c.QueryArray("user")) //获取多个键为user的值

}
func _param(c *gin.Context) {
	fmt.Println(c.Param("user_id")) //除了路径符号以外的所有字符都可以传入
	//http://localhost:8080/param/123465465kjhkjhjk
	//就会把123465465kjhkjhjk打印到后台终端

	fmt.Println(c.Param("order_info"))
	//http://localhost:8080/param/1/2
	//1 2打印到后台终端
}

func _form(c *gin.Context) {
	fmt.Println(c.PostForm("name"))
	fmt.Println(c.PostFormArray("name"))
	fmt.Println(c.DefaultPostForm("age", "18")) //如果没有传自动传入18
	forms, err := c.MultipartForm()             //接收所用的form参数包括文件
	fmt.Println(forms, err)
}
func main() {
	r := gin.Default()
	//r.GET("/query", _query)
	//r.GET("/param/:user_id", _param)
	//r.GET("/param/:user_id/:order_info", _param)
	r.POST("/form", _form)
	r.Run()
}
