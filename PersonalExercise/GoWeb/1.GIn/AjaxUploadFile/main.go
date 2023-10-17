package main

import (
	"github.com/gin-gonic/gin"
	"github.com/my/repo/PersonalExercise/GoWeb/1.GIn/AjaxUploadFile/myFunc"
)

func main() {
	r := gin.Default()
	//写路由：
	//加载html页面：
	r.LoadHTMLGlob("F:\\Users\\gj\\gostudy\\PersonalExercise\\GoWeb\\1.GIn\\AjaxUploadFile\\template/**/*")
	//指定js文件：
	r.Static("/s", "PersonalExercise/GoWeb/1.GIn/AjaxUploadFile/static")
	//定义路由：
	r.GET("/userindex", myFunc.Hello1)
	r.POST("/savefile", myFunc.Hello4)
	r.Run()
}
