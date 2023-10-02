package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// 读取数据文件
	data, err := ioutil.ReadFile("data.txtF:\\Users\\gj\\gostudy\\PythonWork\\1.algorithm\\PythonHomework\\4.work4\\work1\\data.txt")
	if err != nil {
		log.Fatal(err)
	}

	// 将数据按行分割为字符串切片
	lines := strings.Split(string(data), "\n")

	// 过滤数据
	filteredData := make([]string, 0)
	for _, line := range lines {
		if strings.HasPrefix(line, "A") {
			filteredData = append(filteredData, line)
		}
	}

	// 排序数据
	sortedData := make([]string, len(lines))
	copy(sortedData, lines)
	sort.Strings(sortedData)

	// 计数数据
	count := len(lines)

	// 转换数据
	convertedData := make([]int, len(lines))
	for i, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		convertedData[i] = num
	}

	// 打印结果
	fmt.Println("原始数据：", lines)
	fmt.Println("过滤后的数据：", filteredData)
	fmt.Println("排序后的数据：", sortedData)
	fmt.Println("数据数量：", count)
	fmt.Println("转换后的数据：", convertedData)
}
