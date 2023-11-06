package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// handlerFunc
func indexHandler(c *gin.Context) {
	fmt.Println("index")
	name, ok := c.Get("name")
	if !ok {
		name = "匿名用户"
	}
	c.JSON(200, gin.H{
		"msg": name,
	})
	c.JSON(200, gin.H{"msg": "index"})
}

// 定义一个中间件m1
func m1(c *gin.Context) {
	fmt.Println("m1 in..")
	start := time.Now()
	c.Next() //调用后面处理的函数->indexHandler
	//c.Abort() //阻止调用后续的处理函数
	fmt.Println("m1 out..")
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
}

func m2(c *gin.Context) {
	fmt.Println("m2 in..")
	c.Set("name", "ganjing")
	start := time.Now()
	//c.Next()  //调用后面处理的函数->indexHandler
	c.Abort() //阻止调用后续的处理函数
	fmt.Println("m2 out..")
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
}
func AuthMiddleware(doCheck bool) gin.HandlerFunc {
	return func(context *gin.Context) {
		if doCheck {
			context.Next()
		} else {
			context.Next()
		}
	}

}
func main() {
	r := gin.Default()
	r.Use(m1, m2, AuthMiddleware(true)) //全局注册中间件函数m1
	r.GET("/index", indexHandler)
	r.GET("/shop", func(context *gin.Context) { context.JSON(200, gin.H{"msg": "shop"}) })
	r.GET("/user", func(context *gin.Context) { context.JSON(200, gin.H{"msg": "user"}) })

	r.Run()
}
