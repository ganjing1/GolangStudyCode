package main

import (
	"github.com/gin-gonic/gin"
	"github.com/my/repo/PersonalExercise/GoWeb/1.GIn/ajax/myFunc"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("PersonalExercise/GoWeb" +
		"/1.GIn/ajax/template/**/*")
	r.Static("/s", "PersonalExercise/GoWeb/1.GIn/ajax/static")
	r.GET("/userindex", myFunc.Hello1)
	r.POST("/getUserinfo", myFunc.Hello2)
	r.POST("/ajaxpost", myFunc.Hello3)
	r.Run()
}
