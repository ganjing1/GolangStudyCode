package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type charCount struct {
	chCount int
	nuCount int
	spCount int
	otCount int
}

func main() {
	var count charCount
	var content string
	filepath, err := os.OpenFile("F:\\Users\\gj\\gostudy\\PersonalExercise\\basic knowledge\\File\\CalculateChar\\abc.txt", os.O_RDONLY, 066)
	if err != nil {
		fmt.Printf("open have a err:%v\n", err)
		return
	}
	defer filepath.Close()
	reader := bufio.NewReader(filepath)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			content += str
			break
		}
		content += str
	}
	for _, v := range content {
		switch {
		case (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z'):
			count.chCount++
		case (v >= '0' && v <= '9'):
			count.nuCount++
		case (v == ' ' || v == '\t') || (v == '\n'):
			count.spCount++
		default:
			count.otCount++
		}
	}
	fmt.Printf("char have %v,number have %v,space have %v ,other have %v\n", count.chCount, count.nuCount, count.spCount, count.otCount)
}
