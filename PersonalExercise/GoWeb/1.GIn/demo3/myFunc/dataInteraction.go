package myFunc

import "github.com/gin-gonic/gin"

func Hello1(context *gin.Context) {
	id := context.Param("id")
	context.String(200, "获取路径上拼接的参数值 ,%s", id)
}
func Hello2(context *gin.Context) {
	id := context.Param("id")
	context.String(200, "获取路径上拼接的参数值 ,%s", id)
}
func Hello3(context *gin.Context) {
	id := context.Query("id")
	name := context.Query("name")
	context.String(200, "获取路径上拼接的参数值，%s，%s", id, name)
}
func Hello4(context *gin.Context) {
	id := context.DefaultQuery("id", "123")
	name := context.DefaultQuery("name", "甘进")
	context.String(200, "获取路径上拼接的参数值，%s，%s", id, name)
}
func Hello5(context *gin.Context) {
	idvalues := context.QueryArray("id")
	name := context.DefaultQuery("name", "甘进")
	context.String(200, "获取路径上拼接的参数值，%s，%s", idvalues, name)
}
func Hello6(context *gin.Context) {
	user_map := context.QueryMap("user")

	context.String(200, "获取路径上拼接的参数值，%s", user_map)
}
