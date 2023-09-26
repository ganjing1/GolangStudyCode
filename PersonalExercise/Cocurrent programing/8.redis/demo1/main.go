package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
)

func string(ctx context.Context, client *redis.Client) {
	key := "name"
	value := "大脸猫"
	err := client.Set(ctx, key, value, 1*time.Second).Err() //1秒后失效。0表示永不失效
	checkError(err)
	client.Expire(ctx, key, 3*time.Second) //通过Expire设置3秒后失效
	time.Sleep(2 * time.Second)
	v2, err := client.Get(ctx, key).Result()
	checkError(err)
	fmt.Println(v2)
	client.Del(ctx, key)
}

func list(ctx context.Context, client *redis.Client) {
	key := "ids"
	values := []interface{}{1, "中", 3, 4}
	err := client.RPush(ctx, key, values...).Err() //向List右侧插入。如果List不存在会先创建
	checkError(err)
	v2, err := client.LRange(ctx, key, 0, -1).Result() //截取。双闭区间
	checkError(err)
	fmt.Println(v2)
	client.Del(ctx, key)
}

func hashtable(ctx context.Context, client *redis.Client) {
	//key  field1 value1  field2 value2  . ..
	err := client.HSet(ctx, "学生1", "Name", "张三", "Age", 18, "Height", 173.5).Err()
	checkError(err)
	err = client.HSet(ctx, "学生2", "Name", "李四", "Age", 20, "Height", 180.0).Err()
	checkError(err)
	age, err := client.HGet(ctx, "学生2", "Age").Result()
	checkError(err)
	fmt.Println(age)
	for field, value := range client.HGetAll(ctx, "学生1").Val() {
		fmt.Println(field, value)
	}
	client.Del(ctx, "学生1")
	client.Del(ctx, "学生2")
}

func checkError(err error) {
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key不存在")
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", //没有密码
		DB:       0,  //redis默认会创建0-15号DB，这里使用默认的DB
	})
	ctx := context.TODO()
	string(ctx, client)
	list(ctx, client)
	hashtable(ctx, client)
}
