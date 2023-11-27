package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/my/repo/PersonalExercise/GoWeb/QiMi_GoWebProjectProgramming/10.blog_project/logger"
)

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(c *gin.Context) {
		c.String(200, "ok")
	})
	return r
}
