package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
题目描述
输入一行字符，分别统计出包含英文字母、空格、数字和其它字符的个数。

输入描述:
输入一行字符串，可以有空格

输出描述:
统计其中 英文字符，空格字符，数字字符，其他字符的个数

示例1
输入
复制
1qazxsw23 edcvfr45tgbn hy67uj m,ki89ol.\\/;p0-=\\][
输出
复制
26
3
10
12
*/
func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s := input.Text()
		m, n, p, q := static(s)
		fmt.Println(m)
		fmt.Println(n)
		fmt.Println(p)
		fmt.Println(q)
	}
}

func static(s string) (int, int, int, int) {
	m, n, p, q := 0, 0, 0, 0
	for _, v := range s {
		if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' {
			m++
		} else if v == ' ' {
			n++
		} else if v >= '0' && v <= '9' {
			p++
		} else {
			q++
		}
	}
	return m, n, p, q
}
