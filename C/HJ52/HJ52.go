package main

import "fmt"

/*
Levenshtein 距离，又称编辑距离，指的是两个字符串之间，由一个转换成另一个所需的最少编辑操作次数。许可的编辑操作包括将一个字符替换成另一个字符，插入一个字符，删除一个字符。编辑距离的算法是首先由俄国科学家Levenshtein提出的，故又叫Levenshtein Distance。

Ex：

字符串A:abcdefg

字符串B: abcdef

通过增加或是删掉字符”g”的方式达到目的。这两种方案都需要一次操作。把这个操作所需要的次数定义为两个字符串的距离。

要求：

给定任意两个字符串，写出一个算法计算它们的编辑距离。







输入描述:
输入两个字符串

输出描述:
得到计算结果

示例1
输入
复制
abcdefg
abcdef
输出
复制
1
*/

func main() {
	for {
		var s1, s2 string
		n, _ := fmt.Scanln(&s1)
		if n == 0 {
			break
		}
		fmt.Scanln(&s2)

		l1, l2 := len(s1), len(s2)
		arr := make([][]int, l1+1)
		for i := 0; i < l1+1; i++ {
			arr[i] = make([]int, l2+1)
		}
		for i := 0; i < l1+1; i++ {
			arr[i][0] = i
		}
		for i := 0; i < l2+1; i++ {
			arr[0][i] = i
		}
		for i := 1; i < l1+1; i++ {
			for j := 1; j < l2+1; j++ {
				num := 0
				if s1[i-1] != s2[j-1] {
					num = 1
				}
				arr[i][j] = min(arr[i-1][j]+1, arr[i][j-1]+1, arr[i-1][j-1]+num)
			}
		}
		fmt.Println(arr[l1][l2])
	}
}

func min(a, b, c int) int {
	if a > b {
		a = b
	}
	if a > c {
		a = c
	}
	return a
}
