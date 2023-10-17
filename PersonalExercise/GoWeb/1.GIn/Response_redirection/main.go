package main

import (
	"github.com/gin-gonic/gin"
	"github.com/my/repo/PersonalExercise/GoWeb/1.GIn/Response_redirection/myFunc"
)

func main() {
	r := gin.Default()
	//写路由：
	//定义路
	r.GET("/red1", myFunc.Red1)
	r.GET("/red2", myFunc.Red2)
	r.Run()
}
