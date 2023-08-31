package main

import (
	"fmt"
)

type arr [4][4]int

func (t *arr) Print() {
	for _, line := range t {
		for _, number := range line {
			fmt.Printf("%2d ", number)
		}
		fmt.Println()
	}
	fmt.Println()
}

func (t *arr) Left90() {
	temp := new(arr)
	for i, line := range t {
		for j, number := range line {
			temp[len(line)-j-1][i] = number
		}
	}
	*t = *temp
}

func (t *arr) Right90() {
	temp := new(arr)
	for i, line := range t {
		for j, number := range line {
			// fmt.Println(i,j,j,len(line)-1-i)
			temp[j][len(line)-i-1] = number
		}
	}
	*t = *temp
}

func (t *arr) Up() {
	for m := 0; m < 3; m++ {
		for j := 0; j < 4; j++ {
			for i := 0; i < 3; i++ {
				if t[i][j] == 0 {
					t[i][j] = t[i+1][j]
					t[i+1][j] = 0
				} else if t[i][j] == t[i+1][j] {
					t[i][j] = t[i][j] + t[i+1][j]
					t[i+1][j] = 0
				}
			}
		}
	}
}

func main() {
	// t := arr{{0,2,2,4},{0},{},{0,2,0,4}}
	t := arr{{2, 2, 2, 2}, {2, 2, 4, 2}, {}, {0, 2, 0, 4}}
	t.Print()

	//右
	fmt.Println("右")
	t.Left90()
	t.Up()
	t.Right90()
	t.Print()
	//左
	fmt.Println("左")
	t.Right90()
	t.Up()
	t.Left90()
	t.Print()

	//下
	fmt.Println("下")
	t.Right90()
	t.Right90()
	t.Up()
	t.Left90()
	t.Left90()
	t.Print()

	//上
	fmt.Println("上")
	t.Up()
	t.Print()
}
