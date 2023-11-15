package main

import (
	"go.uber.org/zap"
	"net/http"
)

var logger *zap.Logger

func main() {
	InitLogger()
	//1.logger.Sync()是用来将日志缓冲区中的日志写入到输出目的地的操作
	defer logger.Sync()
	//simpleHttpGet("www.google.com")
	simpleHttpGet("http://www.bilibili.com")
}

func InitLogger() {
	//通过调用zap.NewProduction()/zap.NewDevelopment()或者zap.Example()创建一个Logger。
	logger, _ = zap.NewProduction()
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}
