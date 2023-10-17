package main

import "fmt"

func count(s string) int {
	var i int
	i = 0
	for range s {
		i++
	}
	return i

}
func main() {
	fmt.Println(count("hello你好"))
}
