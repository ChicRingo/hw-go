package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*

	参数解析
题目描述
在命令行输入如下命令：

xcopy /s c:\ d:\，

各个参数如下：

参数1：命令字xcopy

参数2：字符串/s

参数3：字符串c:\

参数4: 字符串d:\

请编写一个参数解析程序，实现将命令行各个参数解析出来。



解析规则：

1.参数分隔符为空格
2.对于用“”包含起来的参数，如果中间有空格，不能解析为多个参数。比如在命令行输入xcopy /s “C:\program files” “d:\”时，参数仍然是4个，第3个参数应该是字符串C:\program files，而不是C:\program，注意输出参数时，需要将“”去掉，引号不存在嵌套情况。
3.参数不定长
4.输入由用例保证，不会出现不符合要求的输入




输入描述:
输入一行字符串，可以有空格

输出描述:
输出参数个数，分解后的参数，每个参数都独占一行

示例1
输入
复制
xcopy /s c:\\ d:\\
输出
复制
4
xcopy
/s
c:\\
d:\\
*/
func main() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	out := strings.Fields(str)
	count := 0
	var s []string
	var spaceStr string
	for _, v := range out {
		if v[0] == '"' && v[len(v)-1] == '"' {
			//fmt.Println(v[1 : len(v)-1])
			count++
			s = append(s, v[1:len(v)-1])
		} else if v[0] == '"' && v[len(v)-1] != '"' {
			//fmt.Print(v[1:])
			spaceStr = v[1:] + " "
		} else if v[0] != '"' && v[len(v)-1] == '"' {
			//fmt.Println(v[0 : len(v)-1])
			spaceStr += v[0 : len(v)-1]
			s = append(s, spaceStr)
			count++
		} else {
			//fmt.Println(v)
			s = append(s, v)
			count++
		}
	}
	fmt.Println(count)
	for _, v := range s {
		fmt.Println(v)
	}
}
