package main

import (
	"log"
	"net/http"
	"os"
)

func SetupLogger() {
	//0744表示文件的权限设置为rwxr--r--，即所有者具有读、写和执行权限，但其他用户只有读的权限
	logFileLocation, _ := os.OpenFile("PersonalExercise/GoWeb/QiMi_GoWebProjectProgramming/7.zap/1.logger/test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)

	//log.SetOutput是一个函数，用于设置日志输出的目的地
	log.SetOutput(logFileLocation)
}
func simpleHttpGet(url string) {
	respond, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching url %s:%s", url, err.Error())
	} else {
		log.Printf("status code for %s:%s", url, respond.Status)
		//respond.Body.Close()是用来关闭HTTP响应体的。在使用http.Get方法获取网页内容后，响应体需要被关闭以释放资源
		respond.Body.Close()
	}
}
func main() {
	SetupLogger()
	simpleHttpGet("www.bilibili.com")
	simpleHttpGet(" http://www.google.com")

}
