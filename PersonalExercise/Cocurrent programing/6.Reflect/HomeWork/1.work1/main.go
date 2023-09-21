package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Student struct {
	Name string
}

func structReflect(b interface{}) error {
	v := reflect.ValueOf(b)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return errors.New("传入的参数必须是指针类型")
	}

	v.Elem().FieldByName("Name").SetString("cheerio")
	return nil
}

func main() {
	st := Student{Name: "张三"}
	err := structReflect(&st)
	if err != nil {
		fmt.Println("修改失败:", err)
	} else {
		fmt.Println("修改后的名字为:", st.Name)
	}
}
