package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:hsp@tcp(localhost:3306)/hsp_db02?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Printf("connect to db failed err:%v\n", err)
		return
	}
	fmt.Println("connect success")
}
