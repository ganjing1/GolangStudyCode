package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

//将日志写入文件而不是终端
/*
我们要做的第一个更改是把日志写入文件，而不是打印到应用程序控制台。
我们将使用zap.New(…)方法来手动传递所有配置，
而不是使用像zap.NewProduction()这样的预置方法来创建logger。
*/
func New(core zapcore.Core, options ...Option) *Logger
zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core)
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")
	return zapcore.AddSync(file)
}