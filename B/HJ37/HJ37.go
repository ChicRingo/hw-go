package main

import "fmt"

/*

统计每个月兔子的总数

题目描述
有一只兔子，从出生后第3个月起每个月都生一只兔子，小兔子长到第三个月后每个月又生一只兔子，假如兔子都不死，问每个月的兔子总数为多少？

* 统计出兔子总数。
* @param monthCount 第几个月
* @return 兔子总数

public static int getTotalCount(int monthCount)
{
return 0;
}

本题有多组数据，请使用while (cin>>)读取


输入描述:
输入int型表示month

输出描述:
输出兔子总数int型

示例1
输入
复制
9
输出
复制
34

斐波那切数列

*/

func main() {
	var N int
	for {
		if n, _ := fmt.Scan(&N); n == 0 {
			break
		}
		if N < 1 {
			return
		}
		a, b, c := 1, 0, 0
		for i := 2; i <= N; i++ {
			c += b
			b = a
			a = c
		}
		fmt.Println(a + b + c)
	}
}
