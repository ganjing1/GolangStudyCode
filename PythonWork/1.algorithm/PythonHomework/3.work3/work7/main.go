package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 创建一个哑节点作为合并后链表的头节点
	dummy := &ListNode{}
	// 创建一个指针指向合并后链表的当前节点
	current := dummy
	// 循环比较两个链表的节点值，将较小的节点添加到合并后链表中
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}
	// 将剩余的节点添加到合并后链表的末尾
	if l1 != nil {
		current.Next = l1
	} else if l2 != nil {
		current.Next = l2
	}
	// 返回合并后链表的头节点
	return dummy.Next
}
func main() {
	// 创建两个有序链表
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 4}
	l2 := &ListNode{Val: 1}
	l2.Next = &ListNode{Val: 3}
	l2.Next.Next = &ListNode{Val: 4}
	// 合并两个有序链表
	merged := mergeTwoLists(l1, l2)
	// 打印合并后链表的节点值
	for merged != nil {
		fmt.Println(merged.Val)
		merged = merged.Next
	}
}
