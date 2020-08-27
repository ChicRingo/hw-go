package main

import "fmt"

/*
题目描述
请设计一个算法完成两个超长正整数的加法。


接口说明


/*
请设计一个算法完成两个超长正整数的加法。
输入参数：
String addend：加数
String augend：被加数
返回值：加法结果

本题有多组输入数据，请使用while(cin>>)等方式读取

输入描述:
输入两个字符串数字

输出描述:
输出相加后的结果，string型

示例1
输入
复制
99999999999999999999999999999999999999999999999999
1
输出
复制
100000000000000000000000000000000000000000000000000
*/
func main() {
	for {
		var s1, s2 string
		n, _ := fmt.Scanln(&s1)
		if n == 0 {
			break
		}
		fmt.Scanln(&s2)
		carry := 0
		i := len(s1) - 1
		j := len(s2) - 1
		b := make([]byte, 0)
		for i >= 0 || j >= 0 || carry > 0 {
			sum := carry
			if i >= 0 {
				sum += int(s1[i] - '0')
			}
			if j >= 0 {
				sum += int(s2[j] - '0')
			}
			b = append(b, byte(sum%10)+'0')
			carry = sum / 10
			i--
			j--
		}
		//fmt.Print(b, "        ")
		for i := 0; i < len(b)/2; i++ {
			b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
		}
		fmt.Println(string(b))
	}
}
