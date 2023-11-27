package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func readContent(config *viper.Viper) {
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Printf("解析配置文件出错: %v\n", err)
		}
	}

	// 读取配置
	user1 := config.GetString("section1.user")
	user2 := config.GetString("section2.user")
	height := config.GetInt32("section1.body.height")
	weight := config.GetInt32("section1.body.weight")
	fmt.Println(user1, user2, height, weight)
}

func readJson() {
	config := viper.New()
	config.AddConfigPath("PersonalExercise/GoWeb/QiMi_GoWebProjectProgramming/8.viper_test/conf/") // 文件所在目录
	config.SetConfigName("account")                                                                // 文件名
	config.SetConfigType("json")                                                                   // 文件类型
	readContent(config)
}

func readYaml() {
	config := viper.New()
	config.AddConfigPath("PersonalExercise/GoWeb/QiMi_GoWebProjectProgramming/8.viper_test/conf/") // 文件所在目录
	config.SetConfigName("account")                                                                // 文件名
	config.SetConfigType("yaml")                                                                   // 文件类型

	readContent(config)
}

func main() {
	readJson()
	readYaml()
}
