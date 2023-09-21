package main

import "fmt"

/*返回指定数字在数组中两个数字的和的分别索引*/
func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{} //创建一个哈希表，健和值都是整型
	for i, x := range nums {   //遍历数组
		if p, ok := hashTable[target-x]; ok {
			//p对应target-x,ok是一个布尔值，表示哈希表中是否有target-x这个键
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 17
	res := twoSum(nums, target)
	fmt.Printf("%v", res)
}
