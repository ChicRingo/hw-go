package main

import "fmt"

/*
题目描述
输出7有关数字的个数，包括7的倍数，还有包含7的数字（如17，27，37...70，71，72，73...）的个数（一组测试用例里可能有多组数据，请注意处理）

输入描述:
一个正整数N。(N不大于30000)

输出描述:
不大于N的与7有关的数字个数，例如输入20，与7有关的数字包括7,14,17.

示例1
输入
复制
20
输出
复制
3
*/
func main() {
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			return
		}

		count := 0
		for i := 7; i <= n; i++ {
			if i%7 == 0 || check(i) {
				count++
			}
		}
		fmt.Println(count)
	}
}

func check(x int) bool {
	for x != 0 {
		if x%10 == 7 {
			return true
		}
		x /= 10
	}
	return false
}
