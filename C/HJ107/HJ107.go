package main

import (
	"fmt"
	"math"
)

/*题目描述
•计算一个数字的立方根，不使用库函数

详细描述：

•接口说明

原型：

public static double getCubeRoot(double input)

输入:double 待求解参数

返回值:double  输入参数的立方根，保留一位小数


输入描述:
待求解参数 double类型

输出描述:
输入参数的立方根 也是double类型

示例1
输入
复制
216
输出
复制
6.0
*/
func getCubeRoo(d float64) float64 {
	n := 1.0
	for i := 0.0; i < 1000; {
		s := i * i * i
		if s == d || math.Abs(s-d) < 0.001 {
			return i
		}
		if i*i*i < d {
			i += n
		}
		if i*i*i > d {
			i = i - n
			n = n / 10
			i = i + n
		}
	}
	return 0
}

func main() {
	var num float64
	fmt.Scan(&num)
	fmt.Printf("%.1f\n", getCubeRoo(num))
}
