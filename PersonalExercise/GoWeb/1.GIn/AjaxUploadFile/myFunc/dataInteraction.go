package myFunc

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func Hello1(context *gin.Context) {
	context.HTML(200, "demo1/hello1.html", nil)
}

func Hello2(context *gin.Context) {
	//获取前端传入的文件：
	file, _ := context.FormFile("myfile")

	//加入时间戳
	time_int := time.Now().Unix()
	time_str := strconv.FormatInt(time_int, 10)
	context.SaveUploadedFile(file, "F:\\Users\\gj\\gostudy\\PersonalExercise\\GoWeb\\1.GIn\\AjaxUploadFile\\"+time_str+file.Filename)
	fmt.Println(file.Filename)
}
