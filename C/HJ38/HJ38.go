package main

import "fmt"

/*
题目描述
假设一个球从任意高度自由落下，每次落地后反跳回原高度的一半; 再落下, 求它在第5次落地时，共经历多少米?第5次反弹多高？

最后的误差判断是小数点6位



输入描述:
输入起始高度，int型

输出描述:
分别输出第5次落地时，共经过多少米第5次反弹多高

示例1
输入
复制
1
输出
复制
2.875
0.03125
*/
func main() {
	for {
		var n int
		num, _ := fmt.Scanln(&n)
		if num == 0 {
			break
		}
		f := float64(n)
		sum := f
		temp := f
		for i := 1; i < 5; i++ {
			sum += temp
			temp /= 2
		}
		fmt.Printf("%f\n", sum)
		fmt.Printf("%f\n", temp/2)
	}
}
