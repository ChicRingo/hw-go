package main

import (
	"fmt"
)

/*
字符串反转

题目描述
写出一个程序，接受一个字符串，然后输出该字符串反转后的字符串。（字符串长度不超过1000）

输入描述:
输入N个字符

输出描述:
输出该字符串反转后的字符串

示例1
输入
复制
abcd
输出
复制
dcba
*/

func main() {
	var str string
	fmt.Scanln(&str)

	s := []byte(str)
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	newStr := string(s)
	fmt.Println(newStr)
}
