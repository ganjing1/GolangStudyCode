package main

import (
	"github.com/gin-gonic/gin"
	"github.com/my/repo/PersonalExercise/GoWeb/1.GIn/Template/demo1/myFunc"
	"html/template"
	//"github.com/my/repo/PersonalExercise/GoWeb/1.GIn/AjaxUploadMutiFile/myFunc"
)

func main() {
	r := gin.Default()
	//注册函数:FuncMap是html/funcMap
	r.SetFuncMap(template.FuncMap{
		//键值对的作用：key指定前端调用的名字，value指定的是后端对应的函数
		"add": myFunc.Add,
	})
	r.LoadHTMLGlob("PersonalExercise\\GoWeb\\1.GIn\\Template\\1.demo1\\template/**/*")
	r.GET("/userindex", myFunc.Hello1)
	r.POST("/savefile", myFunc.Hello4)
	r.Run()
}
