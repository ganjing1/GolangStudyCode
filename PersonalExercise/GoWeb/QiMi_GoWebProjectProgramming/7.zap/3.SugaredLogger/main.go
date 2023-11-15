package main

import (
	"go.uber.org/zap"
	"net/http"
)

var sugarLogger *zap.SugaredLogger

func main() {
	InitLogger()
	defer sugarLogger.Sync()
	//simpleHttpGet("www.google.com")
	simpleHttpGet("http://www.souhu.com")
}

func InitLogger() {
	logger, _ := zap.NewProduction()
	sugarLogger = logger.Sugar()
}

func simpleHttpGet(url string) {
	/*
		sugarLogger.Debugf("Trying to hit GET request for %s", url)中的%s是一个占位符，会被url的值所替换。
		这样就可以在日志中打印出类似"Trying to hit GET request for http://www.souhu.com"
	*/
	sugarLogger.Debugf("Trying to hit GET request for %s", url)

	//resp代表一个是 *http.Response 对象，代表对请求的响应
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close() //关闭请求体响应
	}
}
