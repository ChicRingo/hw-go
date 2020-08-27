package main

import (
	"fmt"
)

/*
题目描述
输入一个int型整数，按照从右向左的阅读顺序，返回一个不含重复数字的新的整数。

输入描述:
输入一个int型整数

输出描述:
按照从右向左的阅读顺序，返回一个不含重复数字的新的整数

示例1
输入
复制
9876673
输出
复制
37689
*/
func main() {
	var num string
	for {
		_, err := fmt.Scan(&num)
		if err != nil {
			return
		}
		l := len(num) - 1
		var tab = make(map[string]int, l)
		for i := l; i >= 0; i-- {
			t := string(num[i])
			_, ok := tab[t]
			if !ok {
				tab[t] = 1
				fmt.Print(t)
			}
		}
		fmt.Println()
	}
}
