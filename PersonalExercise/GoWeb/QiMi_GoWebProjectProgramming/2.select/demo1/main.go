package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type user struct {
	id   int
	age  int
	name string
}

var db *sql.DB //声明了一个全局变量 db，类型为 *sql.DB，这意味着它是一个指向 sql.DB 类型的指针。
func queryRowDemo() {
	sqlStr := "2.select id,name,age from u1 where id=?"
	var u user

	//QueryRow是单行查询
	row := db.QueryRow(sqlStr, 1)
	//确保QueryRow之后调用scan方法，否则持有的数据库连接不会被释放
	err := row.Scan(&u.id, &u.name, &u.age) //在sqlstr的查询语句中找id，name，和age并存放到u中对应的id，name，age中
	if err != nil {
		fmt.Printf("scan failed, err: %v\n", err)
		return
	}
	fmt.Printf("id: %d name: %s age: %d\n", u.id, u.name, u.age)

	//Query是多行查询
	sqlStr2 := "2.select name,id,age from u1 where age!=?"
	rows, err := db.Query(sqlStr2, 1)
	if err != nil {
		panic(err)
		return
	}
	for rows.Next() {
		var u user
		err := rows.Scan(&u.name, &u.id, &u.age) //select里面的字段次序要和scan里面的一致
		if err != nil {
			panic(err)

		}
		fmt.Println(u.id, u.name, u.age)
	}
}

func insertRowDemo() {
	sqlStr := "insert into u1(name,age) values (?,?)"
	ret, err := db.Exec(sqlStr, "xiaogan", 22)
	if err != nil {
		fmt.Printf("insert faild ,err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("获取id失败 ,err:%v\n", err)
		return
	}
	fmt.Printf("insert success ,the id is %d.\n", theID)
}

func updateRowDemo() {
	ud := "update u1 set name=? where age=?"
	result, err := db.Exec(ud, "甘进", 22)
	if err != nil {
		panic(err)
		return
	}
	n, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("修改受影响的行数为  ", n)
}

func deleteRowDemo() {
	ud := "delete from u1  where age=?"
	result, err := db.Exec(ud, 22)
	if err != nil {
		panic(err)
		return
	}
	n, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("删除受影响的行数为  ", n)
}

// 预处理
func prepareQueryDemo() {
	sqlStr := "2.select id, name, age from u1 where id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

func initMySQL() (err error) {
	dsn := "root:hsp@tcp(localhost:3306)/hsp_db02"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("connect to db failed, err: %v\n", err)
		return
	}
	db.SetConnMaxLifetime(time.Second * 10)
	db.SetMaxOpenConns(200) //最大连接
	db.SetMaxIdleConns(10)  //最大空闲连接
	return
}

func main() {
	if err := initMySQL(); err != nil {
		fmt.Printf("connect to db failed ,err:%v\n", err)
	}
	fmt.Println("connect tp db success")
	//queryRowDemo()
	//insertRowDemo()
	//updateRowDemo()
	//deleteRowDemo()
	prepareQueryDemo()

	/*
		总结：增删改都要用都要调用db.exec
		单行查询则是用queryrow

		多行查询用db.query,多行查询会返回一个多列的结果我们用rows接收，
		然后调用if rows.Next()判断是否存在某列
	*/
}
