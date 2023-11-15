package myFunc

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func Hello1(context *gin.Context) {
	context.HTML(200, "1.demo1/hello1.html", nil)
}

func Hello4(context *gin.Context) {
	//获取前端传入的文件：
	form, _ := context.MultipartForm()
	files, _ := form.File["myfile"] //File是个Map，通过key获取value部分
	for _, file := range files {
		//加入一个时间戳：
		time_int := time.Now().Unix()
		time_str := strconv.FormatInt(time_int, 10) //10:十进制
		//保存在我的本地：
		context.SaveUploadedFile(file, "PersonalExercise\\GoWeb\\1.GIn\\1.demo1"+time_str+file.Filename)
	}
}
