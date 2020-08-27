package main

import (
	"fmt"
	"strings"
)

/*
公共字串计算
题目描述
题目标题：

计算两个字符串的最大公共字串的长度，字符不区分大小写

详细描述：

接口说明

原型：

int getCommonStrLength(char * pFirstStr, char * pSecondStr);

输入参数：

     char * pFirstStr //第一个字符串

     char * pSecondStr//第二个字符串


输入描述:
输入两个字符串

输出描述:
输出一个整数

示例1
输入
复制
asdfas
werasdfaswer
输出
复制
6
*/
func main() {
	var s1, s2 string
	for {
		n, err := fmt.Scan(&s1, &s2)
		if n < 2 || err != nil {
			return
		}

		max := 0
		for i := 0; i < len(s1); i++ {
			for j := i + 1; j <= len(s1); j++ {
				if strings.Contains(s2, s1[i:j]) && j-i > max {
					max = j - i
				}
			}
		}
		fmt.Println(max)
	}
}
