package main

import "fmt"

/*
题目描述
输入一个int型的正整数，计算出该int型数据在内存中存储时1的个数。

输入描述:
 输入一个整数（int类型）

输出描述:
 这个数转换成2进制后，输出1的个数

示例1
输入
复制
5
输出
复制
2d
*/
func main() {
	var num int
	fmt.Scanf("%d", &num)
	count := 0
	for {
		if num == 0 {
			break
		}
		if num%2 == 1 {
			count++
		}
		num /= 2
	}
	fmt.Println(count)
}
