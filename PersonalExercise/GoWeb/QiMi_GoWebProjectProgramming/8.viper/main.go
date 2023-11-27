package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath("PersonalExercise/GoWeb/QiMi_GoWebProjectProgramming/8.viper/") // 还可以在工作目录中查找配置
	viper.SetConfigName("config")                                                       // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")                                                         // 如果配置文件的名称中没有扩展名，则需要配置此项
	//viper.SetConfigFile("config.yaml")
	//viper.AddConfigPath("/etc/appname/")  // 查找配置文件所在的路径
	//viper.AddConfigPath("$HOME/.appname") // 多次调用以添加多个搜索路径

	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		//配置文件发生变更后会调用的回调函数
		fmt.Println("config file changed", e.Name)
	})
	r := gin.Default()
	r.GET("version", func(c *gin.Context) {
		c.String(200, viper.GetString("version"))
	})
	r.Run()
}
