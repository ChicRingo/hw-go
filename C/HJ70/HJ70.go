package main

import "fmt"

/*
题目描述
矩阵乘法的运算量与矩阵乘法的顺序强相关。


例如：

    A是一个50×10的矩阵，B是10×20的矩阵，C是20×5的矩阵

计算A*B*C有两种顺序：（（AB）C）或者（A（BC）），前者需要计算15000次乘法，后者只需要3500次。


编写程序计算不同的计算顺序需要进行的乘法次数


输入描述:
输入多行，先输入要计算乘法的矩阵个数n，每个矩阵的行数，列数，总共2n的数，最后输入要计算的法则

输出描述:
输出需要进行的乘法次数

示例1
输入
复制
3
50 10
10 20
20 5
(A(BC))
输出
复制
3500
*/
func main() {
	for {
		n := 0
		if _, err := fmt.Scan(&n); err != nil {
			return
		}
		nums := make([][2]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&nums[i][0], &nums[i][1])
		}

		s := ""
		fmt.Scan(&s)
		var res [][]int
		ans, idx := 0, 0

		for _, v := range s {
			if v == ')' && len(res) >= 2 {
				num1 := res[len(res)-1]
				num2 := res[len(res)-2]
				res = res[0 : len(res)-2]
				ans += num2[0] * num2[1] * num1[1]
				num3 := []int{num2[0], num1[1]}
				res = append(res, num3)
			} else if v >= 'A' && v <= 'Z' {
				num3 := []int{nums[idx][0], nums[idx][1]}
				res = append(res, num3)
				idx++
			}
		}

		fmt.Println(ans)
	}
}
