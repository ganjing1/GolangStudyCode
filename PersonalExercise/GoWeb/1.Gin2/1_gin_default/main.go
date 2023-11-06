package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ganjin(c *gin.Context) {
	c.String(200, "ganjing")
}

/*
通过gin框架在本地前端网页显示hello world
*/
func main() {
	r := gin.Default()
	r.GET("/", Ganjin)
	//gin会把web服务运行在本机的0.0.0.0：8080端口上
	//让本机的所有ip都能访问
	//r.Run("0.0.0.0:8080")

	//原生http服务的方式
	http.ListenAndServe("localhost:8080", r)
}
