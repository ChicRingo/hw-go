package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
题目描述
将一个字符串str的内容颠倒过来，并输出。str的长度不超过100个字符。 如：输入“I am a student”，输出“tneduts a ma I”。

输入参数：
inputString：输入的字符串

返回值：
输出转换好的逆序字符串

输入描述:
输入一个字符串，可以有空格

输出描述:
输出逆序的字符串

示例1
输入
复制
I am a student
输出
复制
tneduts a ma I
*/
func main() {
	inReader := bufio.NewReader(os.Stdin)
	s, _ := inReader.ReadSlice('\n')
	// 去掉末尾空格 10
	s = s[:len(s)-1]
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	fmt.Println(string(s))
}
