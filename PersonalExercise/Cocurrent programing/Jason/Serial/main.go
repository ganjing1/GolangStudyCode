package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string
	Age      int
	Birthday string
	Salary   float64
	Skill    string
}

func testStruct() {
	ganjing := Monster{
		Age:      22,
		Name:     "甘进",
		Birthday: "2001/01/10",
		Salary:   3001.5,
		Skill:    "Golang Mysql",
	}
	//将ganjing 序列化
	//json.Marshal()函数用于将数据序列化为JSON格式。它的参数是一个指向要序列化的数据的指针
	data, err := json.Marshal(&ganjing)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("ganjing序列化后=%v\n", string(data)) //data本身是字节切片，所以要用string才能表示内容
}
func testMap() {
	var a map[string]interface{} //a这个map的key是字符串，value是任意类型
	//使用map需要，make
	a = make(map[string]interface{})
	a["name"] = "张三"
	a["age"] = 22
	a["sex"] = "man"
	/*
		a本身就是一个引用类型，所以我们直接将a传递给json.Marshal()函数，
		而不需要使用取地址操作符&。这是因为a已经是一个指针，指向了map类型的数据
	*/
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("map a序列化后:%v\n", string(data))
}

// 对切片进行序列化,我们这个切片[]map[string]interface{}
func testSlice() {
	var slice []map[string]interface{} //切片类型，其中每个元素都是一个 map[string]interface{} 类型的值
	var m1 map[string]interface{}      //map类型
	var m2 map[string]interface{}
	/*
		map 是一种无序的键值对集合，其中的键和值可以是任意类型。
		map[string]interface{} 表示键的类型是字符串，
		值的类型是 interface{}，即可以接收任意类型的值。
	*/

	//使用map前，需要先make
	m1 = make(map[string]interface{})
	m2 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "北京"
	m2["name"] = "marry"
	m2["age"] = "22"
	m2["address"] = "纽约"

	slice = append(slice, m1)
	slice = append(slice, m2)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("Slice1 marshal have an err=%v", err)
		return
	}

	fmt.Printf("序列化后的silce：%v ", string(data))
}
func main() {
	//将结构体，map，切片进行序列化
	testStruct()
	testMap()
	testSlice()
}
