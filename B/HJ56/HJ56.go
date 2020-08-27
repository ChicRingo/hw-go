package main

import "fmt"

/*

题目描述

完全数（Perfect number），又称完美数或完备数，是一些特殊的自然数。

它所有的真因子（即除了自身以外的约数）的和（即因子函数），恰好等于它本身。

例如：28，它有约数1、2、4、7、14、28，除去它本身28外，其余5个数相加，1+2+4+7+14=28。

给定函数count(int n),用于计算n以内(含n)完全数的个数。计算范围, 0 < n <= 500000

返回n以内完全数的个数。 异常情况返回-1

 *  给定函数count(int n),用于计算n以内(含n)完全数的个数

 * @param n  计算范围, 0 < n <= 500000

 * @return n 以内完全数的个数, 异常情况返回-1

输入描述:
输入一个数字

输出描述:
输出完全数的个数
*/

func test(num int) bool {
	if num <= 1 {
		return false
	}
	var sum int
	for i := 1; i <= num/2; i++ {
		if num%i == 0 {
			sum += i
		}
	}
	return sum == num
}

func main() {
	var num int
	for {
		res, _ := fmt.Scanln(&num)
		if res == 0 || num == 0 {
			break
		}
		var pf int
		for i := 1; i <= num; i++ {
			if test(i) {
				pf += 1
			}
		}
		fmt.Println(pf)
	}
}

func main1() {
	for {
		var n int
		fmt.Scanf("%d", &n)
		if n <= 0 {
			break
		}
		var res int
		if n < 6 {
			res = 0
		} else if n >= 6 && n < 28 {
			res = 1
		} else if n >= 28 && n < 496 {
			res = 2
		} else if n >= 496 && n < 8128 {
			res = 3
		} else if n >= 8128 && n < 130816 {
			res = 4
		} else {
			res = 5
		}
		fmt.Println(res)
	}

}
