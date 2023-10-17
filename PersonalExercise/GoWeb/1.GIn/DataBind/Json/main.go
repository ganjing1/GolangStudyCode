package main

import (
	"github.com/gin-gonic/gin"
	"github.com/my/repo/PersonalExercise/GoWeb/1.GIn/DataBind/Json/myfunc"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("PersonalExercise\\GoWeb\\1.GIn\\DataBind\\Json\\templates/**/*")
	r.Static("/s", "PersonalExercise\\GoWeb\\1.GIn\\DataBind\\Json\\static")
	r.GET("/userindex", myfunc.Hello)
	r.POST("/toajax", myfunc.Hello2)
	r.Run()
}
