package main

import "fmt"

func main() {
	var a int64
	fmt.Scanln(&a)
	if a >= 90 && a <= 100 {
		fmt.Println("A")
	} else if a >= 80 && a < 90 {
		fmt.Println("B")
	} else if a >= 70 && a < 80 {
		fmt.Println("C")
	} else if a >= 60 && a < 70 {
		fmt.Println("D")
	} else  {
		fmt.Println("pass")

	}
}

