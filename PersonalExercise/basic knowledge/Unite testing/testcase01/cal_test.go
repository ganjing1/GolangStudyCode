package main

import (
	"fmt"
	"testing" //引入go的testing包
)

func TestAdd(t *testing.T) { //必须要以Test为开头,且Test后面的字母非a-z，那么测试自动调用
	res := add(10)
	if res != 55 {
		//fmt.Printf("Add(10)，错误 。期望值：%v，实际值：%v\n",55,res)
		t.Fatalf("Add(10)，错误 。期望值：%v，实际值：%v\\n", 55, res)
	}
	t.Logf("Add(10)执行正确...")
}

func TestHello(t *testing.T) {
	fmt.Println("this is TestHello")
}
func TestSub(t *testing.T) { //必须要以Test为开头
	res := getSub(10, 20)
	if res != -10 {
		//fmt.Printf("Add(10)，错误 。期望值：%v，实际值：%v\n",55,res)
		t.Fatalf("sub(10,20)，错误 。期望值：%v，实际值：%v\\n", -10, res)
	}
	t.Logf("Sub(10,20)执行正确...")
}
