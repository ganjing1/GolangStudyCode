package main

import "fmt"

func main() {
	//byte uint8的别名 ascii码
	//rune int32的别名
	var c1 byte = 'c'
	var c2 rune = 'c'
	fmt.Println(c1, c2)
	fmt.Println("c1:%T  c2:%T\n", c1, c2)
	s := "hello甘进"
	for i := 0; i < len(s); i++ {
		fmt.Println("%c", s[i])
	}
	for _, r := range s {
		fmt.Printf("%c\n", r)
	}
}
