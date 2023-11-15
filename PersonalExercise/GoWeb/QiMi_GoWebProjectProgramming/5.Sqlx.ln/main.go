package main

import (
	"database/sql/driver"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strings"
)

var DB *sqlx.DB

func initDB() (err error) {
	dsn := "root:hsp@tcp(localhost:3306)/hsp_db02?charset=utf8mb4&parseTime=True&loc=Local"
	// 也可以使用MustConnect连接不成功就panic
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(10)
	return
}

type User struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}

// BatchInsertUsers 自行构造批量插入的语句
func BatchInsertUsers(users []*User) error {
	// 存放 (?, ?) 的slice
	valueStrings := make([]string, 0, len(users))
	// 存放values的slice
	valueArgs := make([]interface{}, 0, len(users)*2)
	// 遍历users准备相关数据
	for _, u := range users {
		// 此处占位符要与插入值的个数对应
		valueStrings = append(valueStrings, "(?, ?)")
		valueArgs = append(valueArgs, u.Name)
		valueArgs = append(valueArgs, u.Age)
	}
	// 自行拼接要执行的具体语句
	stmt := fmt.Sprintf("INSERT INTO user (name, age) VALUES %s",
		strings.Join(valueStrings, ","))
	_, err := DB.Exec(stmt, valueArgs...)
	return err
}
func (u User) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}

// BatchInsertUsers2 使用sqlx.In帮我们拼接语句和参数, 注意传入的参数是[]interface{}
func BatchInsertUsers2(users []interface{}) error {
	query, args, _ := sqlx.In("insert into user(name,age)values (?),(?),(?)", users...)
	fmt.Println(query)
	fmt.Println(args)
	_, err := DB.Exec(query, args...)
	if err != nil {
		panic(err)
	}
	return err
}

// BatchInsertUsers3 使用NamedExec实现批量插入
func BatchInsertUsers3(users []*User) error {
	_, err := DB.NamedExec("INSERT INTO user (name, age) VALUES (:name, :age)", users)
	return err
}
func QueryByIDs(ids []int) (users []User, err error) {
	query, args, err := sqlx.In("select * from user where id in (?)", ids)
	if err != nil {
		panic(err)
	}
	//sqlx.in 返回带 ？ bindvar 的查询语句 ，我们使用Rebind()重新绑定它
	query = DB.Rebind(query)
	DB.Select(&users, query, args...)
	return
}
func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	defer DB.Close()
	//u1 := User{Name: "aaa", Age: 18}
	//u2 := User{Name: "bbb", Age: 28}
	//u3 := User{Name: "ccc", Age: 38}

	//// 方法1
	//users := []*User{&u1, &u2, &u3}
	//err = BatchInsertUsers(users)
	//if err != nil {
	//	fmt.Printf("BatchInsertUsers failed, err:%v\n", err)
	//}

	// 方法2
	//user2 := []interface{}{u1, u2, u3}
	//BatchInsertUsers2(user2)

	//
	//// 方法3
	//users3 := []*User{&u1, &u2, &u3}
	//err = BatchInsertUsers3(users3)
	//if err != nil {
	//	fmt.Printf("BatchInsertUsers3 failed, err:%v\n", err)

	users, err := QueryByIDs([]int{1, 2, 3, 4, 5, 6})
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		fmt.Printf("users:%v\n", user)
	}

}
