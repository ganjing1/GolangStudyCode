package main

import (
	"github.com/gin-gonic/gin"
	"github.com/my/repo/PersonalExercise/GoWeb/1.GIn/demo3/myFunc"
)

func main() {
	r := gin.Default()
	r.GET("/demo/:id", myFunc.Hello1)
	r.GET("/demo2/*id", myFunc.Hello2)
	r.GET("/demo3", myFunc.Hello3)
	r.GET("/demo4", myFunc.Hello4)
	r.GET("/demo5", myFunc.Hello5)
	r.GET("/demo6", myFunc.Hello6)
	r.Run()
}
