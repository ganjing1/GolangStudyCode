package main

import (
	"github.com/gin-gonic/gin"
	"github.com/my/repo/PersonalExercise/GoWeb/1.GIn/post/myFunc"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("PersonalExercise/GoWeb/1.GIn/post/template/**/*")
	r.GET("/userindex", myFunc.Hello1)
	r.POST("/getUserinfo", myFunc.Hello2)
	r.Run()
}
