package main

import (
	"fmt"
	"sort"
)

/*
题目描述
给出一个名字，该名字有26个字符串组成，定义这个字符串的“漂亮度”是其所有字母“漂亮度”的总和。
每个字母都有一个“漂亮度”，范围在1到26之间。没有任何两个字母拥有相同的“漂亮度”。字母忽略大小写。
给出多个名字，计算每个名字最大可能的“漂亮度”。
输入描述:
整数N，后续N个名字

输出描述:
每个名称可能的最大漂亮程度

示例1
输入
复制
2
zhangsan
lisi
输出
复制
192
101
*/
func main() {
	var n int
	for {
		num, _ := fmt.Scan(&n)
		if num == 0 {
			break
		}
		arr := make([]string, n)
		for i := 0; i < n; i++ {
			var temp string
			fmt.Scan(&temp)
			arr[i] = temp
		}

		for _, v := range arr {
			fmt.Println(test(v))
		}

	}
}

func test(s string) int {
	arr := make([]int, 26)
	for _, v := range s {
		arr[v-'a']++
	}
	sort.Ints(arr)
	sum := 0
	for i := 0; i < 26; i++ {
		sum += arr[i] * (i + 1)
	}
	return sum
}
