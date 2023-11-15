package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 100,
	})
	_, err = rdb.Ping().Result()
	return err
}

func redisExample() {
	err := rdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed ,err:%v\n", err)
		return
	}
	val, err := rdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed ,err:%v\n", err)
		return
	}
	fmt.Println("score", val)

	val2, err := rdb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed ,err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}
}

func hgetDemo() {
	v, err := rdb.HGetAll("user").Result()
	if err != nil {
		//	redis.nil
		//其他错误
		fmt.Printf("hgetall failed ,err:%v\n", err)
		return
	}
	fmt.Println(v)
	v2 := rdb.HMGet("user", "name", "age").Val()
	fmt.Println(v2)
	v3 := rdb.HGet("user", "age").Val()
	fmt.Println(v3)
}

// zsetDemo 操作zset示例

func main() {
	err := initClient()
	if err != nil {
		fmt.Printf("init redis client failed ，err:%v\n", err)
		return
	}
	fmt.Println("connect redis success")
	defer rdb.Close() //释放相关资源
	//redisExample()
	//hgetDemo()

}
