package main

import (
	"github.com/gin-gonic/gin"
	"github.com/my/repo/PersonalExercise/GoWeb/1.GIn/DataBind/Routine/myfunc"
)

func main() {
	r := gin.Default()
	r.GET("/userindex", myfunc.Hello)
	r.Run()
}
