package work1

import (
	"encoding/json"
	"fmt"
	"os"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (this *Monster) Store() bool {
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
		return false
	}
	//保存到文件
	filePath := "./monster.ser"
	err = os.WriteFile(filePath, data, 0666) //将data写入到filePath对应的路径文件下
	if err != nil {
		fmt.Println("write file err=", err)
		return false
	}
	return true
}
func (this *Monster) ReStore() bool {
	filePath := "./monster.ser"
	data, err := os.ReadFile(filePath) //filePath中的内容技术序列化后的
	if err != nil {
		fmt.Println("readfile err=", err)
		return false
	}
	err = json.Unmarshal(data, this) //data表示要反序列化的内容，this，表示要存放解码结果的变量
	if err != nil {
		fmt.Println("Unmarshal err=", err) //反序列化错误
		return false
	}
	return true
}
