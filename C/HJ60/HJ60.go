package main

import "fmt"

/*
题目描述
任意一个偶数（大于2）都可以由2个素数组成，组成偶数的2个素数有很多种情况，本题目要求输出组成指定偶数的两个素数差值最小的素数对
输入描述:
输入一个偶数

输出描述:
输出两个素数

示例1
输入
复制
20
输出
复制
7
13
*/
func demo(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	for {
		var num, a int
		_, err := fmt.Scan(&num)
		if err != nil {
			return
		}
		a = num >> 1
		for t := a; t < num; t++ {
			if demo(t) && demo(num-t) {
				fmt.Println(num - t)
				fmt.Println(t)
				break
			}
		}
	}
}
