package main

import (
	"github.com/gin-gonic/gin"
	"github.com/my/repo/PersonalExercise/GoWeb/1.GIn/Routine_group/demo1/myfunc"
)

func main() {
	r := gin.Default()
	//加载html页面：
	r.LoadHTMLGlob("PersonalExercise\\GoWeb\\1.GIn\\Routine_group\\templates/**/*")
	//指定文件：
	r.Static("/s", "part11/static")
	//按照版本号对路由进行分组：
	v1 := r.Group("/version01")
	{
		v1.GET("/userindex", myfunc.Hello1)
		v1.GET("/toFormBind", myfunc.Hello2)
		v1.GET("/userindex2", myfunc.Hello3)
	}
	v2 := r.Group("/version02")
	{
		v2.GET("/userindex3", myfunc.Hello4)
		v2.POST("/toajax", myfunc.Hello5)
		v2.GET("/userindex4/:uname/:age", myfunc.Hello6)
		v2.GET("/userindex4/丽丽/18", myfunc.Hello1)
	}
	r.Run()
}
