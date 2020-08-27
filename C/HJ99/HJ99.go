package main

import "fmt"

/*
题目描述
自守数是指一个数的平方的尾数等于该数自身的自然数。例如：25^2 = 625，76^2 = 5776，9376^2 = 87909376。请求出n以内的自守数的个数


接口说明


/*
功能: 求出n以内的自守数的个数


输入参数：
int n

返回值：
n以内自守数的数量。

本题有多组输入数据，请使用while(cin>>)等方式处理


输入描述:
int型整数

输出描述:
n以内自守数的数量。

示例1
输入
复制
2000
输出
复制
8
*/
func main() {
	for {
		var num int
		_, err := fmt.Scan(&num)
		if err != nil {
			break
		}

		count := 0
		var flag int
		for i := 0; i <= num; i++ {
			flag = 10

			for {
				if i/flag == 0 {
					break
				}
				flag = flag * 10
			}

			if (i*i)%flag == i {
				count++
			}
		}
		fmt.Println(count)
	}
}
