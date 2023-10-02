package main

import (
	"github.com/gin-gonic/gin"
	"github.com/my/repo/PersonalExercise/GoWeb/1.GIn/demo2/myFunc"
)

func main() {
	r := gin.Default()
	//r.LoadHTMLFiles("template/hello01.html", "template/hello2.html")

	/*会加载所有匹配该路径的模板文件，并将它们存储在内存中，以便后续使用*/
	//template/**/*下面有几级就用**代表目录级别，例如访问三层级别下的template/**/**/*
	r.LoadHTMLGlob("F:\\Users\\gj\\gostudy\\PersonalExercise\\GoWeb\\1.GIn\\demo2\\template/**/*")

	/*客户端可以通过访问 /s/xxx 的 URL 来获取 static 目录下的 xxx 文件*/
	r.Static("/s", "F:\\Users\\gj\\gostudy\\PersonalExercise\\GoWeb\\1.GIn\\demo2\\static")

	/*调用了 r.GET 方法，传入了两个参数 “/demo1” 和 Hello，
	这个方法会将第一个参数指定的 URL 路径和第二个参数指定的处理函数绑定起来，也就是说，
	当客户端发送 GET 请求到 /demo1 时，会执行 Hello 函数，并返回相应的 HTML 内容。*/
	r.GET("/demo1", myFunc.Hello)       //对字符串进行渲染
	r.GET("/demo2", myFunc.Hello2)      //对结构体进行渲染
	r.GET("/demo3", myFunc.Array)       //对数组进行渲染
	r.GET("/demo4", myFunc.StructArray) //对结构体数组进行渲染
	r.GET("/demo5", myFunc.Map)         //对map进行渲染
	r.GET("/demo6", myFunc.MapStruct)   //对MapStruct进行渲染
	r.GET("/demo7", myFunc.Slice)       //对Slice进行渲染
	r.Run("127.0.0.1:8080")             //默认是本机端口
}
