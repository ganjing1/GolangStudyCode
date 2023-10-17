package main

import (
	"github.com/gin-gonic/gin"
	"github.com/my/repo/PersonalExercise/GoWeb/1.GIn/DataBind/Form/myfunc"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("PersonalExercise\\GoWeb\\1.GIn\\DataBind\\Form\\templates/**/*")

	r.GET("/userindex", myfunc.Hello1)
	r.GET("/toForm", myfunc.Hello2)
	r.Run()
}
