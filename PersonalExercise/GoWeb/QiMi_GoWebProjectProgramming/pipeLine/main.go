package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

// 创建一个redis客户端
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

// 定义一个阈值
const maxValue = 10

// 定义一个事务函数
func txf(tx *redis.Tx) error {
	// 获取键的当前值或者0
	n, err := tx.Get(context.Background(), "mykey").Int()
	if err != nil && err != redis.Nil {
		return err
	}

	// 如果值小于阈值，就创建一个pipeline
	if n < maxValue {
		_, err = tx.TxPipelined(context.Background(), func(pipe redis.Pipeliner) error {
			// 在pipeline中将键的值加一
			pipe.Incr(context.Background(), "mykey")
			return nil
		})
	}
	return err
}

func main() {
	// 使用Watch方法执行事务函数，监视键的变化
	err := rdb.Watch(context.Background(), txf, "mykey")
	if err != nil {
		panic(err)
	}

	// 打印键的最终值
	val, err := rdb.Get(context.Background(), "mykey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("mykey:", val)
}
