package main

import (
	"github.com/gin-gonic/gin"
	"github.com/my/repo/PersonalExercise/GoWeb/1.GIn/DataBind/URL/myfunc"
)

func main() {
	r := gin.Default()
	//加载html页面：
	r.LoadHTMLGlob("PersonalExercise\\GoWeb\\1.GIn\\DataBind\\URL\\templates/**/*")

	r.GET("/userindex4/:uname/:age", myfunc.Hello6)
	r.Run()
}
