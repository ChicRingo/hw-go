package main

import "fmt"

/*
题目描述
功能:输入一个正整数，按照从小到大的顺序输出它的所有质因子（重复的也要列举）（如180的质因子为2 2 3 3 5 ）

最后一个数后面也要有空格

输入描述:
输入一个long型整数

输出描述:
按照从小到大的顺序输出它的所有质数的因子，以空格隔开。最后一个数后面也要有空格。

示例1
输入
复制
180
输出
复制
2 2 3 3 5
*/

func main() {
	for {
		var i, num, tmp uint64
		_, err := fmt.Scan(&num)
		if err != nil {
			return
		}
		tmp = num
		for i = 2; i <= tmp; i++ {
			if num == 1 {
				break
			}
			if i == tmp {
				fmt.Printf("%d%s", i, " ")
				break
			}
			for num%i == 0 {
				num /= i
				fmt.Printf("%d%s", i, " ")
			}
		}
		fmt.Println()
	}
}
