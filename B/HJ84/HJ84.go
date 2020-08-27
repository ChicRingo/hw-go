package main

import "fmt"

/*
题目描述
找出给定字符串中大写字符(即'A'-'Z')的个数

接口说明

原型：int CalcCapital(String str);

返回值：int



输入描述:
输入一个String数据

输出描述:
输出string中大写字母的个数

示例1
输入
复制
add123#$%#%#O
输出
复制
1
*/

func main() {
	for {
		s := ""
		_, err := fmt.Scan(&s)
		if err != nil {
			return
		}
		count := 0
		for i := 0; i < len(s); i++ {
			if s[i:i+1] >= "A" && s[i:i+1] <= "Z" {
				count++
			}
		}
		fmt.Println(count)
	}
}
