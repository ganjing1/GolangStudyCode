package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func startMysql() error {
	dsn := "root:hsp@tcp(localhost:3306)/hsp_db02"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("连接数据库失败!")
		panic(err)

	}

	return err

}

func transaction() {
	ts, err := db.Begin()
	if err != nil {
		panic(err)
		ts.Rollback()
	}

	sqlStr := "update u1 set age=21 where id=?"
	rt, err := db.Exec(sqlStr, 2)
	affectNum, _ := rt.RowsAffected()
	fmt.Println("更新事务的相关条数：", affectNum)
	sqlStr1 := "insert into u1(name,age) values (? ,?)"
	rt1, err := db.Exec(sqlStr1, "zzzz", 100)
	if err != nil {
		fmt.Println("插入事务执行失败")
		ts.Rollback()
	}
	affectNum1, _ := rt1.RowsAffected()
	fmt.Println("插入事务的相关条数：", affectNum1)
	if affectNum1 != 0 && affectNum != 0 {
		ts.Commit()
		fmt.Println("更新和插入事务执行成功！")
	}
}
func main() {
	err := startMysql()
	if err != nil {
		fmt.Printf("connect to db failed ,err:%v\n", err)
	} else {
		fmt.Println("连接成功！")
	}
	transaction()
}
