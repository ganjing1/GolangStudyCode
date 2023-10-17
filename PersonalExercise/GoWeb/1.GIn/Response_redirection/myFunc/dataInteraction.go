package myFunc

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Red1(context *gin.Context) {
	fmt.Println("this is Red1")
	context.Redirect(http.StatusFound, "/red2")
}

func Red2(context *gin.Context) {
	fmt.Println("succes Redirect ")
	context.String(http.StatusOK, "这是Red2")
}
