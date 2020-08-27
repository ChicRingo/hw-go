package main

import "fmt"

/*
题目描述
请编写一个函数（允许增加子函数），计算n x m的棋盘格子（n为横向的格子数，m为竖向的格子数）沿着各自边缘线从左上角走到右下角，总共有多少种走法，要求不能走回头路，即：只能往右和往下走，不能往左和往上走。
输入描述:
输入两个正整数

输出描述:
返回结果

示例1
输入
复制
2
2
输出
复制
6
*/

func main() {
	var n, m int
	for {
		if _, err := fmt.Scan(&n, &m); err != nil {
			return
		}
		fmt.Println(sum(m, n))
	}
}
func sum(x, y int) int {
	if x == 0 || y == 0 {
		return 1
	}
	return sum(x-1, y) + sum(x, y-1)
}

func main1() {
	var m, n int
	for {
		if _, err := fmt.Scan(&m, &n); err != nil {
			return
		}
		a, b := 1, 1
		for i := m + n; i >= n+1; i-- {
			a *= i
		}
		for j := m; j >= 1; j-- {
			b *= j
		}
		fmt.Println(a / b)
	}
}
