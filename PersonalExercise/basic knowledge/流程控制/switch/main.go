package main

import "fmt"

func  main() {
	//第一种
	var score int64
	fmt.Scanln(&score)
	switch score/10{
	case 10:{fmt.Println("A")};
	case 9:{fmt.Println("A")};
	case 8:{fmt.Println("B")};
	case 7:fmt.Println("C")
			fallthrough//fallthrough如果执行case 7 那么执行case7
			// 下面的case6 只执行一次
	case 6:{fmt.Println("D")};
	default:{fmt.Println("pass")}
	}

	//第二种：判断不写在switch后面
	//var score int64
	//fmt.Scanln(&score)
	//c := score / 10//
	//switch {
	//case c == 10:
	//	{
	//		fmt.Println("A")
	//	};
	//case c == 9:
	//	{
	//		fmt.Println("A")
	//	};
	//case c == 8:
	//	{
	//		fmt.Println("B")
	//	};
	//case c == 7:
	//	{
	//		fmt.Println("C")
	//	};
	//case c == 6:
	//	{
	//		fmt.Println("D")
	//	};
	//default:
	//	{
	//		fmt.Println("pass")
	//	}

	//第三种
	//这种 switch 后面要跟分号
	//var score int
	//switch c := score / 10; {
	//case c >= 5:
	//	{
	//		fmt.Println("A")
	//	};
	//case c < 5:
	//	{
	//		fmt.Println("B")
	//	};
	//default:
	//	{
	//		fmt.Println("pass")
	//	}
	//
	//}
}
