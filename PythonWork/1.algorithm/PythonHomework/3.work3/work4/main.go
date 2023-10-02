package main

import "fmt"

func multiplyMatrix(matrix1 [][]int, matrix2 [][]int) [][]int {
	// 获取矩阵1的行数和列数
	rows1 := len(matrix1)
	cols1 := len(matrix1[0])
	// 获取矩阵2的行数和列数
	//rows2 := len(matrix2)
	cols2 := len(matrix2[0])
	// 创建结果矩阵
	result := make([][]int, rows1)
	for i := 0; i < rows1; i++ {
		result[i] = make([]int, cols2)
	}
	// 计算乘积
	for i := 0; i < rows1; i++ {
		for j := 0; j < cols2; j++ {
			for k := 0; k < cols1; k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}
	return result
}
func main() {
	matrix1 := [][]int{{1, 2, 3}, {4, 5, 6}}
	matrix2 := [][]int{{7, 8}, {9, 10}, {11, 12}}
	result := multiplyMatrix(matrix1, matrix2)
	fmt.Println("结果矩阵：")
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[0]); j++ {
			fmt.Printf("%d ", result[i][j])
		}
		fmt.Println()
	}
}
