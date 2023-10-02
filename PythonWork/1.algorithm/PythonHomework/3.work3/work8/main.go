package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	// 使用快慢指针法找到链表的中间节点
	slow := head
	fast := head
	// 快指针每次移动两步，慢指针每次移动一步，当快指针到达链表末尾时，慢指针正好指向中间节点
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 返回中间节点
	return slow
}
func main() {
	// 创建一个链表
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	// 计算链表的中间节点
	middle := middleNode(head)
	// 打印中间节点的值
	fmt.Println(middle.Val)
}
