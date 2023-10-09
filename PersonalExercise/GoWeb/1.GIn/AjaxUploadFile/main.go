package main

import (
	"github.com/gin-gonic/gin"
	"github.com/my/repo/PersonalExercise/GoWeb/1.GIn/AjaxUploadFile/myFunc"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("PersonalExercise\\GoWeb\\1.GIn\\AjaxUploadFile\\template//**/*")
	r.Static("/s", "PersonalExercise\\GoWeb\\1.GIn\\AjaxUploadFile\\static")
	r.GET("/userindex", myFunc.Hello1)
	r.POST("/savefile", myFunc.Hello2)
	r.Run()
}
