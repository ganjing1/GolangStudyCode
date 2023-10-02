package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	// 使用快慢指针法判断链表是否有环
	slow := head
	fast := head
	// 快指针每次移动两步，慢指针每次移动一步，如果存在环，快指针最终会追上慢指针
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		// 如果快指针追上慢指针，说明链表有环
		if slow == fast {
			return true
		}
	}
	// 如果快指针到达链表末尾，说明链表无环
	return false
}
func main() {
	// 创建一个链表
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	// 判断链表是否有环
	hasCycle := hasCycle(head)
	// 打印判断结果
	fmt.Println(hasCycle)
}
