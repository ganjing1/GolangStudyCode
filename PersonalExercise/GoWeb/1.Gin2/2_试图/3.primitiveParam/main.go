package main

import (
	"github.com/gin-gonic/gin"
)

func _raw(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		panic(err)
	}
	println(string(body))
}
func main() {
	r := gin.Default()
	r.POST("/raw", _raw)
	r.Run()
}
