package main

import "fmt"

func main() {
	var chessMap [11][11]int
	chessMap[1][2] = 1           //黑子
	chessMap[2][3] = 2           //蓝子
	for _, v := range chessMap { //外层遍历行
		for _, v2 := range v { //内层遍历列
			fmt.Printf("%d\t", v2)
		}
		fmt.Println("\n")
	}
	//转成稀疏数组。
	/*

	 */
}
