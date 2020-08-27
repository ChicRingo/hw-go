package main

import (
	"fmt"
)

/*
题目描述
正整数A和正整数B 的最小公倍数是指 能被A和B整除的最小的正整数值，设计一个算法，求输入A和B的最小公倍数。

输入描述:
输入两个正整数A和B。

输出描述:
输出A和B的最小公倍数。

示例1
输入
复制
5 7
输出
复制
35
*/
func main() {
	var a, b int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		return
	}
	fmt.Println(a * b / min(a, b))
}
func min(a, b int) int {
	if a < b {
		a, b = b, a
	}
	c := a % b
	if c == 0 {
		//b是最小公约数
		return b
	}
	//否则继续循环
	return min(b, c)
}
