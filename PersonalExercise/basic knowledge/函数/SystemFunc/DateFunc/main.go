package  main

import (
	"fmt"
	"time"
)

func main(){
	now:=time.Now()
	fmt.Printf("%v,\n%T",now,now)

	fmt.Printf("year:%v\n",now.Year())
	fmt.Printf("month:%v\n",int(now.Month()))
	fmt.Printf("day:%v\n",now.Day())
	fmt.Printf("Minute:%v\n",now.Minute())
	fmt.Printf("year:%v\n",now.Second())
	fmt.Println()
	datestr:=fmt.Sprintf("%d-%d-%d   %d:%d:%d",now.Year(),now.Month(),
		now.Day(),now.Hour(),now.Minute(),now.Second())
	fmt.Println(datestr)

	//这个参数字符串中的各个数字必须是固定的，必须这样写
	datestr2:=now.Format("2006/01/02 15/04/01")
	fmt.Println(datestr2)

	datestr3:=now.Format("2006 16:04")
	fmt.Println(datestr3)
}
