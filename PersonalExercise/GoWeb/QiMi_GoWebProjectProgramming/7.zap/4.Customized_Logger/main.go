//将日志写入文件而不是终端
/*
我们要做的第一个更改是把日志写入文件，而不是打印到应用程序控制台。
我们将使用zap.New(…)方法来手动传递所有配置，
而不是使用像zap.NewProduction()这样的预置方法来创建logger。
*/

package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
)

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

func getEncoder() zapcore.Encoder {
	//已jason的形式存储日志

	encoderConfig := zapcore.EncoderConfig{
		EncodeTime: zapcore.ISO8601TimeEncoder,
	}
	return zapcore.NewConsoleEncoder(encoderConfig)

}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.OpenFile("PersonalExercise/GoWeb/QiMi_GoWebProjectProgramming/"+
		"7.zap/4.Customized_Logger/test.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0744)

	//函数接受一个io.Writer接口类型的参数，并返回一个对应的zapcore.WriteSyncer类型的对象，该对象可以用于将日志写入到指定的目标。
	return zapcore.AddSync(file)
	//zapcore.WriteSyncer是一个接口，用于表示可以同步写入日志的对象。
}

func InitLogger() {
	//调用 getLogWriter 函数，该函数返回一个将日志写入文件的对象，并将其赋值给 writeSyncer 变量。
	writeSyncer := getLogWriter()
	encoder := getEncoder()

	//core使用了JSON编码器和文件输出目标，因此创建的日志记录器将使用JSON格式编码日志，并将日志写入到文件中。
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core)

	//通过调用logger.Sugar()，可以创建一个新的基于sugar的日志记录器对象，该记录器提供了一些更方便的方法来记录日志，例如Debugw、Infow、Errorw等，
	//这些方法可以用来记录结构化的日志，同时也可以添加上下文信息和键值对字段。
	sugarLogger = logger.Sugar()
}

func simpleHttpGet(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}

func main() {
	InitLogger()
	defer sugarLogger.Sync()
	simpleHttpGet("http://www.baidu.com")
}

/*
这段代码是一个使用Uber 的 Zap 日志库的示例，它展示了如何将日志输出到文件而不是终端。让我来解释一下整个执行流程：

InitLogger 函数初始化了日志记录器。它使用 getLogWriter 函数获取一个将日志写入文件的对象，使用 getEncoder 函数获取一个 JSON 编码器，并创建一个新的核心日志记录器，其中包含编码器和写入器。

getEncoder 函数返回一个 JSON 编码器，它使用了 zap.NewProductionEncoderConfig 来配置编码器。

getLogWriter 函数创建一个文件，用于将日志写入其中。在这里，它创建了一个名为 test.log 的文件，并将其用作日志的输出目标。

simpleHttpGet 函数使用 sugarLogger 记录了一些日志信息。在这个例子中，它记录了尝试进行 GET 请求和请求结果的信息。

在 main 函数中，首先调用 InitLogger 来初始化日志记录器。然后使用 defer 关键字在程序结束时执行 sugarLogger.Sync()，以确保所有日志都被写入文件。最后，调用 simpleHttpGet 函数来尝试进行 HTTP GET 请求。

在执行程序时，它将尝试发送一个 HTTP GET 请求到百度网站，并记录请求和响应的相关信息到日志文件中。
*/
