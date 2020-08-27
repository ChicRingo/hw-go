package main

import "fmt"

/*
题目描述
功能: 求一个byte数字对应的二进制数字中1的最大连续数，例如3的二进制为00000011，最大连续2个1

输入: 一个byte型的数字

输出: 无

返回: 对应的二进制数字中1的最大连续数
输入描述:
输入一个byte数字

输出描述:
输出转成二进制之后连续1的个数

示例1
输入
复制
3
输出
复制
2
*/
func main() {
	var n byte
	for {
		if _, err := fmt.Scan(&n); err != nil {
			return
		}
		var k int
		for n != 0 {
			k++
			n = n & (n << 1)
		}
		fmt.Println(k)
	}
}
