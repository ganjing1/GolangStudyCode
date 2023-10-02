package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/testdemo", func(context *gin.Context) {
		context.String(200, "这是我的第一个Gin")
	})
	r.Run() //默认是本机端口
}
