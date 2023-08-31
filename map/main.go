package main
import "fmt"
import "math/rand"
import "sort"
func main() {
	/* //map（映射）
	//光声明map类型，但是没有初始化 a就是初始值nil
	var a map[string]int
	fmt.Println(a==nil)
	//map的初始化
	a=make(map[string]int, 8)
	fmt.Println(a==nil)
	//map中如何添加键值对
	a["甘进"]=100
	a["TIDY"]=200
	fmt.Printf("a:%#v\n",a)
	fmt.Printf("type:%T\n",a)
	//声明map的同时完成初始化
	b:=map[int]bool{
		1: true,
		2: false,
	}
	fmt.Printf("b:%#v\n",b)
	fmt.Printf("type:%T\n",b) */
	//判断某个键存不存在
	 /*var scoremap=make(map[string]int,8)
	scoremap["甘进"]=100//["甘进"]是键，100是值，组成键值对
	scoremap["Tidy"]=200
	//判断lol在不在map中
	/*
	value,ok:=scoremap["Tidy"]
	if ok{
		fmt.Println("lol在scoremap中",value)//如果ok的值在scoremap那么执行此条
	}else{
		fmt.Println("lol不在scoremap中！")//如果ok的值不在scoremap那么执行此条
	}*/
	//遍历map for range
	//map是无序的，键值对和添加的顺序无关，map是无序的，键值对和添加的顺序无关
	/* for k,v:=range scoremap{
		fmt.Println(k,v)
		//只遍历map中的key
		for k:=range scoremap{
			fmt.Println(k)
		}
		//只遍历map中的value
		for _,V:=range scoremap{
			fmt.Println(V)
		}
	}
	//删除 甘进 这个键值对
	delete(scoremap,"甘进")
	fmt.Println(scoremap)
 */		
 var scoremap=make(map[string]int,10)
//添加50个键值对
for i:=0;i<50;i++{
	key:= fmt.Sprintf("stu%02d",i)
	value:= rand.Intn(100)//0-99的随机数
	scoremap[key]=value
} 
for k,v:=range scoremap{//前面的k是键，v是值
	fmt.Println(k,v) 
}
//按照key从小到大的顺序去遍历scoremap
//1.先取出所有的key存放到切片中
keys:=make([]string,0,100)
for k:=range scoremap{
	keys = append(keys,k)
}
//2.对key做排列
sort.Strings(keys)//keys目前是一个有序的片段
//3.按照排列后的key对scoremap排序
for _,key:=range keys{
	fmt.Println(key,scoremap[key])
}
} 