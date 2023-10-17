package myFunc

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func Hello1(context *gin.Context) {
	context.HTML(200, "demo1/hello1.html", nil)
}

func Hello4(context *gin.Context) {
	//获取前端传入的文件：
	file, _ := context.FormFile("myfile")

	//加入时间戳
	time_int := time.Now().Unix()
	time_str := strconv.FormatInt(time_int, 10)
	context.SaveUploadedFile(file, "PersonalExercise\\GoWeb\\1.GIn\\AjaxUploadFile\\"+time_str+file.Filename)
	context.String(200, "文件上传成功")
}
