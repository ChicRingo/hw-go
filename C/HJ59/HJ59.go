package main

import "fmt"

/*
题目描述
找出字符串中第一个只出现一次的字符

输入描述:
输入几个非空字符串

输出描述:
输出第一个只出现一次的字符，如果不存在输出-1

示例1
输入
复制
asdfasdfo
aabb
输出
复制
o
-1
*/
func main() {
	for {
		var str string
		var a bool
		_, err := fmt.Scan(&str)
		if err != nil {
			return
		}
		m := make(map[rune]int)
		for _, b := range str {
			m[b]++
		}
		for _, b := range str {
			if m[b] == 1 {
				fmt.Println(string(b))
				a = true
				break
			}
		}
		if !a {
			fmt.Println(-1)
		}
	}
}
