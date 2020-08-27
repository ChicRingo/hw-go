package main

import (
	"fmt"
)

/*
题目描述
题目描述

把M个同样的苹果放在N个同样的盘子里，允许有的盘子空着不放，问共有多少种不同的分法？（用K表示）5，1，1和1，5，1 是同一种分法。


输入

每个用例包含二个整数M和N。0<=m<=10，1<=n<=10。


样例输入

7 3


样例输出

8


/**

* 计算放苹果方法数目


* 输入值非法时返回-1

* 1 <= m,n <= 10

* @param m 苹果数目

* @param n 盘子数目数

* @return 放置方法总数

*

*/

//func main() {
//	for {
//		var m, n int
//		if length, _ := fmt.Scan(&m, &n); length == 0 {
//			break
//		}
//		dp := make([][][]int, m+1)
//		for i := 0; i <= m; i++ {
//			dp[i] = make([][]int, n+1)
//			for j := 0; j <= n; j++ {
//				dp[i][j] = make([]int, m+1)
//			}
//		}
//		for j := 0; j <= n; j++ {
//			dp[0][j][0] = 1
//		}
//		for i := 1; i <= m; i++ {
//			for j := 0; j <= n; j++ {
//				for k := 0; k <= m; k++ {
//					for a := 0; a <= j && a*i <= k; a++ {
//						dp[i][j][k] += dp[i-1][j-a][k-i*a]
//					}
//				}
//			}
//		}
//		fmt.Println(dp[m][n][m])
//	}
//}

func CalCnt(m, n int) int {
	if m < 0 || n < 0 {
		return 0
	}
	if n == 1 || m == 1 {
		return 1
	}
	return CalCnt(m, n-1) + CalCnt(m-n, n)
	//1.假设有一个盘子为空，则(m,n)问题转化为将m个苹果放在n-1个盘子上，即求得(m,n-1)即可
	//2.假设所有盘子都装有苹果，则每个盘子上至少有一个苹果，即最多剩下m-n个苹果，问题转化为将m-n个苹果放到n个盘子上
	//即求(m-n，n)
	//7 3
	//72+43
	//71+52+42+13
	//1+51+32+1+22+12+1+
}
func main() {
	for {
		var m, n int
		if _, err := fmt.Scan(&m, &n); err != nil {
			break
		}
		fmt.Println(CalCnt(m, n))
	}
}
