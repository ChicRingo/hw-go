package main

import "fmt"

/*
密码提取
题目描述
Catcher 是MCA国的情报员，他工作时发现敌国会用一些对称的密码进行通信，比如像这些ABBA，ABA，A，123321，但是他们有时会在开始或结束时加入一些无关的字符以防止别国破解。比如进行下列变化 ABBA->12ABBA,ABA->ABAKK,123321->51233214　。因为截获的串太长了，而且存在多种可能的情况（abaaab可看作是aba,或baaab的加密形式），Cathcer的工作量实在是太大了，他只能向电脑高手求助，你能帮Catcher找出最长的有效密码串吗？

（注意：记得加上while处理多个测试用例）

输入描述:
输入一个字符串

输出描述:
返回有效密码串的最大长度

示例1
输入
复制
ABBA
输出
复制
4
*/

func main() {
	var s string
	for {
		if _, err := fmt.Scan(&s); err != nil {
			return
		}

		var s1 string
		for i := len(s) - 1; i >= 0; i-- {
			s1 += string(s[i])
		}

		max := 0
		dp := make([]int, len(s))

		for j := 0; j < len(s); j++ {
			for i := len(s) - 1; i >= 0; i-- {
				if s[j] == s1[i] {
					if i >= 1 {
						dp[i] = dp[i-1] + 1
					} else {
						dp[i] = 1
					}
					if max < dp[i] {
						max = dp[i]
					}
				} else {
					dp[i] = 0
				}
			}
		}

		fmt.Println(max)

	}
}
