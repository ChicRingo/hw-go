package main

import "fmt"

/*
题目描述
题目说明

蛇形矩阵是由1开始的自然数依次排列成的一个矩阵上三角形。

样例输入

5

样例输出

1 3 6 10 15

2 5 9 14

4 8 13

7 12

11

输入描述:
输入正整数N（N不大于100）

输出描述:
输出一个N行的蛇形矩阵。

示例1
输入
复制
4
输出
复制
1 3 6 10
2 5 9
4 8
7
*/
func main() {
	var n int

	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		// 上一行开始的数
		lrow := 1
		// 当前打印的数
		rrow := 0
		for i := 1; i <= n; i++ {
			lrow += i - 1
			//fmt.Println("row: ", lrow)
			rrow = lrow
			for j := 1; j <= n-i+1; j++ {
				if j == n-i+1 {
					// 最后一个
					fmt.Printf("%d", rrow)
				} else {
					fmt.Printf("%d ", rrow)
				}
				rrow += j + i
			}
			fmt.Println()
		}
	}
}
