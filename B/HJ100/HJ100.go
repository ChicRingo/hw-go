package main

import "fmt"

/*
题目描述
功能:等差数列 2，5，8，11，14。。。。

输入:正整数N >0

输出:求等差数列前N项和

返回:转换成功返回 0 ,非法输入与异常返回-1

本题为多组输入，请使用while(cin>>)等形式读取数据

输入描述:
输入一个正整数。

输出描述:
输出一个相加后的整数。

示例1
输入
复制
2
输出
复制
7
*/

func main() {
	var a int
	for {
		_, err := fmt.Scan(&a)
		if err != nil {
			return
		}
		all := 0
		for i := 1; i <= a; i++ {
			all = all + (3*i - 1)
		}
		fmt.Println(all)
		//公式f(n) = (1 + 3*n) * n / 2
		//fmt.Println((3*a*a + a) >> 1)
	}
}
