package main

import (
	"container/list"
	"fmt"
)

type Student struct {
	Name   string
	Scores map[string]int
}
type Class struct {
	Name     string
	Students *list.List
}

func main() {
	class := &Class{
		Name:     "Class 1",
		Students: list.New(),
	}
	// 添加学生
	student1 := &Student{
		Name:   "Alice",
		Scores: map[string]int{"Math": 80, "English": 90},
	}
	class.Students.PushBack(student1)
	student2 := &Student{
		Name:   "Bob",
		Scores: map[string]int{"Math": 70, "English": 85},
	}
	class.Students.PushBack(student2)
	// 统计每个同学的成绩
	for e := class.Students.Front(); e != nil; e = e.Next() {
		student := e.Value.(*Student)
		totalScore := 0
		for _, score := range student.Scores {
			totalScore += score
		}
		averageScore := totalScore / len(student.Scores)
		fmt.Printf("%s: Total score: %d, Average score: %d\n", student.Name, totalScore, averageScore)
	}
	// 统计每门课程的成绩
	courseScores := make(map[string]int)
	for e := class.Students.Front(); e != nil; e = e.Next() {
		student := e.Value.(*Student)
		for course, score := range student.Scores {
			courseScores[course] += score
		}
	}
	for course, score := range courseScores {
		fmt.Printf("%s: Total score: %d\n", course, score)
	}
}
