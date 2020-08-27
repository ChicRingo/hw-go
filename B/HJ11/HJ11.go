package main

import (
	"fmt"
	"strconv"
)

/*
数字颠倒

题目描述
描述：

输入一个整数，将这个整数以字符串的形式逆序输出

程序不考虑负数的情况，若数字含有0，则逆序形式也含有0，如输入为100，则输出为001

输入描述:
输入一个int整数

输出描述:
将这个整数以字符串的形式逆序输出

示例1
输入
复制
1516000
输出
复制
0006151
*/
func main() {
	var num int
	fmt.Scanln(&num)

	str := test(num)
	fmt.Println(str)
}
func test(num int) (str string) {
	for {
		x := num % 10
		if x == 0 {
			break
		}
		str = str + strconv.Itoa(x)
		fmt.Println(x)
		num = num / 10
	}
	return
}
