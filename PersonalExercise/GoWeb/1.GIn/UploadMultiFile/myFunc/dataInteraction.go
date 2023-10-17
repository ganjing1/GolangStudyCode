package myFunc

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func Hello1(context *gin.Context) {
	context.HTML(200, "demo1/hello1.html", nil)
}

func Hello2(context *gin.Context) {
	form, _ := context.MultipartForm()
	files := form.File["myfile"]
	for _, file := range files {
		time_int := time.Now().Unix()
		time_str := strconv.FormatInt(time_int, 10)
		context.SaveUploadedFile(file, "F:\\Users\\gj\\gostudy\\PersonalExercise\\GoWeb\\1.GIn\\UploadMultiFile/"+file.Filename+time_str)
	}
}
