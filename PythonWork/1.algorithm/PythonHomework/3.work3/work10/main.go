package main

import "fmt"

type Stack struct {
	data []int
}

func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
}
func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic("Stack is empty")
	}
	length := len(s.data)
	val := s.data[length-1]
	s.data = s.data[:length-1]
	return val
}
func (s *Stack) Top() int {
	if s.IsEmpty() {
		panic("Stack is empty")
	}
	return s.data[len(s.data)-1]
}
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}
func main() {
	// 创建一个栈
	stack := &Stack{}
	// 入栈
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	// 获取栈顶元素
	top := stack.Top()
	fmt.Println("Top:", top)
	// 出栈
	val1 := stack.Pop()
	val2 := stack.Pop()
	val3 := stack.Pop()
	fmt.Println("Pop:", val1, val2, val3)
	// 判断栈是否为空
	empty := stack.IsEmpty()
	fmt.Println("IsEmpty:", empty)
}
