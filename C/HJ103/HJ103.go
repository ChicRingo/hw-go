package main

import "fmt"

/*
题目描述
题目描述

   Redraiment是走梅花桩的高手。Redraiment总是起点不限，从前到后，往高的桩子走，但走的步数最多，不知道为什么？你能替Redraiment研究他最多走的步数吗？



样例输入

6

2 5 1 5 4 5



样例输出

3



提示

Example:
6个点的高度各为 2 5 1 5 4 5
如从第1格开始走,最多为3步, 2 4 5
从第2格开始走,最多只有1步,5
而从第3格开始走最多有3步,1 4 5
从第5格开始走最多有2步,4 5

所以这个结果是3。
接口说明

方法原型：

    int GetResult(int num, int[] pInput, List  pResult);

输入参数：
   int num：整数，表示数组元素的个数（保证有效）。
   int[] pInput: 数组，存放输入的数字。

输出参数：
   List pResult: 保证传入一个空的List，要求把结果放入第一个位置。
返回值：
  正确返回1，错误返回0

输入描述:
输入多行，先输入数组的个数，再输入相应个数的整数

输出描述:
输出结果

示例1
输入
复制
6
2
5
1
5
4
5
输出
复制
3
*/
var num int

func main() {
	for {
		if _, err := fmt.Scan(&num); err != nil {
			return
		}
		m := make([]int, num)
		var n int
		for i := 0; i < num; i++ {
			fmt.Scan(&n)
			m[i] = n
		}
		res := make([]int, num)
		var ans int
		for i := num - 1; i >= 0; i-- {
			for j := i + 1; j < num; j++ {
				if m[i] < m[j] && res[i] < res[j]+1 {
					res[i] = res[j] + 1
				}
			}
			if res[i] > ans {
				ans = res[i]
			}
		}
		fmt.Println(ans + 1)
	}
}
