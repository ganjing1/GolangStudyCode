package main

import (
	"encoding/json"
	"fmt"
)

func UnMarshalStruct() {
	//str 是通过网络传输获取到的，
	str := "{\"monster_name\":\"甘进\",\"monster_age\":22,\"monster_birthday\":" +
		"\"2001/01/10\",\"monster_salary\":3001.5,\"r_skill\":\"Golang Mysql\"}\n"
	var monster Monster

	//Unmarshal接收两个参数
	//1.是字符串所对应的Byte切片
	//2.是把反序列化的str交给谁，要用引用传递
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("Unmarshal failed err=%v\n", err)
	}

	fmt.Printf("反序列化后的字符串为:%v\n", monster)
	fmt.Printf("monster_name为:%v\n", monster.Name)

}

func UnmarshalMap() {
	str := "{\"age\":22,\"name\":\"张三\",\"sex\":\"man\"}\n"
	// 定义map
	var student map[string]interface{} //student这个map的健是string类型 值是任意类型

	//反序列化map不需要make map 因为make map操作被封装到Nnmarshal 操作里面去了
	//创建student map
	//student = make(map[string]interface{})

	err := json.Unmarshal([]byte(str), &student) //当判断为map类型时自动make了
	if err != nil {
		fmt.Printf("UnmarshalMap have a err:%v\n", err)
	}
	fmt.Printf("unmarshalMap 反序列化后：%v\n", student)
}

func UnmarshalSlice() {
	str := "[{\"address\":\"北京\", \"age\":\"7\", \"name\":\"jack\"}," +
		"{\"address\":\"上海\", \"age\":\"11\", \"name\":\"marry\"}]\n"
	//定义一个切片类型
	var slice []map[string]interface{}
	//同样不需要make
	//定义map

	err := json.Unmarshal([]byte(str), &slice) //当判断为map类型时自动make了
	if err != nil {
		fmt.Printf("UnmarshalSlice have a err:%v\n", err)
	}
	fmt.Printf("unmarshalSlice 反序列化后：%v\n", slice)
}

type Monster struct {
	Name     string  `json:"monster_name"`
	Age      int     `json:"monster_age"`
	Birthday string  `json:"monster_birthday"`
	Salary   float64 `json:"monster_salary"`
	Skill    string  `json:"monster_skill"`
}

func main() {
	UnMarshalStruct()
	UnmarshalMap()
	UnmarshalSlice()
}
