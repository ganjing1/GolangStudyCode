package main

import (
	"github.com/gin-gonic/gin"
	"github.com/my/repo/PersonalExercise/GoWeb/1.GIn/Template1/demo1/myFunc"
	//"github.com/my/repo/PersonalExercise/GoWeb/1.GIn/AjaxUploadMutiFile/myFunc"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("PersonalExercise\\GoWeb\\1.GIn\\Template1\\1.demo1\\template/**/*")
	r.Static("/s", "PersonalExercise\\GoWeb\\1.GIn\\Template1\\1.demo1\\static")

	r.GET("/userindex3", myFunc.Hello4)
	r.POST("/toajax", myFunc.Hello5)
	r.Run()
}
