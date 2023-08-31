package db

import "fmt"


func GetConn(){//首字母大写就是共有的，小写私有的
	fmt.Println("db包下的getConn ")
}

