package main

import "fmt"

/*
题目描述
将一个字符中所有出现的数字前后加上符号“*”，其他字符保持不变
public static String MarkNum(String pInStr)
{

return null;
}
注意：输入数据可能有多行
输入描述:
输入一个字符串

输出描述:
字符中所有出现的数字前后加上符号“*”，其他字符保持不变

示例1
输入
复制
Jkdi234klowe90a3
输出
复制
Jkdi*234*klowe*90*a*3*
*/

func main() {
	for {
		var str string
		_, err := fmt.Scan(&str)
		if err != nil {
			return
		}
		str1 := ""
		for i := 0; i < len(str); i++ {
			if str[i] < '0' || str[i] > '9' {
				str1 += string(str[i])
			} else {
				if i == 0 || (str[i-1] < '0' || str[i-1] > '9') {
					str1 += "*"
				}
				str1 += string(str[i])
				if i == len(str)-1 || (str[i+1] < '0' || str[i+1] > '9') {
					str1 += "*"
				}
			}
		}
		fmt.Println(str1)
	}
}
