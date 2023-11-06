package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "GET",
		})
		r.POST("/index", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"method": "POST",
			})
		})

		r.DELETE("/index", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"method": "DELETE",
			})
		})

		r.PUT("/index", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"method": "PUT",
			})
		})
		r.Any("/user", func(c *gin.Context) {
			switch c.Request.Method {
			case "GET":
				c.JSON(200, gin.H{"method": "GET"})
			case "POST":
				c.JSON(200, gin.H{"method": "POST"})
			case "http.MethodPost":
				c.JSON(http.StatusOK, gin.H{"method": "POST"})

			}

			c.JSON(200, gin.H{
				"method": "Any",
			})
		})
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"msg": "ganjin.com"})
	})
	//r.GET("/video/index", func(c *gin.Context) {
	//	c.JSON(200, gin.H{"msg": "/video/index"})
	//})
	//r.GET("/shop/index", func(c *gin.Context) {
	//	c.JSON(200, gin.H{"msg": "/shop/index"})
	//})

	//路由组设置:公用的前缀提取出来，创建一个路由

	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(context *gin.Context) {
			{
				context.JSON(200, gin.H{"msg": "index"})
			}
		})
		videoGroup.GET("/image", func(context *gin.Context) {
			{
				context.JSON(200, gin.H{"msg": "image"})
			}
		})
	}

	r.Run()
}
