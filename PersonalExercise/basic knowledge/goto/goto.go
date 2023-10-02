package main

import "fmt"

func main() {
	var o bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				//设置退出标签
				o = true
				break
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		//分层for循环判断
		if o {
			break
		}
	}
}
