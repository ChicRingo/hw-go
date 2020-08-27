package main

import "fmt"

/*
题目描述
密码按如下规则进行计分，并根据不同的得分为密码进行安全等级划分。

一、密码长度:

5 分: 小于等于4 个字符

10 分: 5 到7 字符

25 分: 大于等于8 个字符

二、字母:

0 分: 没有字母

10 分: 全都是小（大）写字母

20 分: 大小写混合字母

三、数字:

0 分: 没有数字

10 分: 1 个数字

20 分: 大于1 个数字

四、符号:

0 分: 没有符号

10 分: 1 个符号

25 分: 大于1 个符号

五、奖励:

2 分: 字母和数字

3 分: 字母、数字和符号

5 分: 大小写字母、数字和符号

最后的评分标准:

>= 90: 非常安全

>= 80: 安全（Secure）

>= 70: 非常强

>= 60: 强（Strong）

>= 50: 一般（Average）

>= 25: 弱（Weak）

>= 0:  非常弱


对应输出为：

VERY_SECURE

SECURE,

VERY_STRONG,

STRONG,

AVERAGE,

WEAK,

VERY_WEAK,


请根据输入的密码字符串，进行安全评定。

注：

字母：a-z, A-Z

数字：-9

符号包含如下： (ASCII码表可以在UltraEdit的菜单view->ASCII Table查看)

!"#$%&'()*+,-./     (ASCII码：x21~0x2F)

:;<=>?@             (ASCII<=><=><=><=><=>码：x3A~0x40)

[\]^_`              (ASCII码：x5B~0x60)

{|}~                (ASCII码：x7B~0x7E)

接口描述：

Input Param
String pPasswordStr:    密码，以字符串方式存放。

Return Value
根据规则评定的安全等级。

输入描述:
输入一个string的密码

输出描述:
输出密码等级

示例1
输入
复制
38$@NoNoNo
输出
复制
VERY_SECURE
*/
func main() {
	var s string
	for {
		if _, err := fmt.Scan(&s); err != nil {
			return
		}

		res := 0
		l := len(s)
		if l <= 4 {
			res += 5
		} else if l <= 7 {
			res += 10
		} else {
			res += 25
		}

		lower, upper, digit, sym := 0, 0, 0, 0
		for _, v := range s {
			if v >= 'a' && v <= 'z' {
				lower++
			} else if v >= 'A' && v <= 'Z' {
				upper++
			} else if v >= '0' && v <= '9' {
				digit++
			} else {
				sym++
			}
		}
		fmt.Println(lower, upper, digit, sym)
		//0 分: 没有字母
		//10 分: 全都是小（大）写字母
		//20 分: 大小写混合字母
		if lower != 0 && upper == 0 || lower == 0 && upper != 0 {
			res += 10
		} else if lower != 0 && upper != 0 {
			res += 20
		}
		//0 分: 没有数字
		//10 分: 1 个数字
		//20 分: 大于1 个数字
		if digit == 1 {
			res += 10
		} else if digit > 1 {
			res += 20
		}
		//0 分: 没有符号
		//10 分: 1 个符号
		//25 分: 大于1 个符号
		if sym == 1 {
			res += 10
		} else if sym > 1 {
			res += 25
		}
		//奖励:
		//2 分: 字母和数字
		//3 分: 字母、数字和符号
		//5 分: 大小写字母、数字和符号
		if digit != 0 { //有数字
			if sym == 0 { //无符号
				if lower != 0 || upper != 0 {
					res += 2
				}
			} else { //有符号
				if lower == 0 && upper != 0 || lower != 0 && upper == 0 {
					res += 3
				} else if lower != 0 && upper != 0 {
					res += 5
				}
			}
		}

		result := "VERY_WEAK"
		if res >= 90 {
			result = "VERY_SECURE"
		} else if res >= 80 {
			result = "SECURE"
		} else if res >= 70 {
			result = "VERY_STRONG"
		} else if res >= 60 {
			result = "STRONG"
		} else if res >= 50 {
			result = "AVERAGE"
		} else if res >= 25 {
			result = "WEAK"
		}
		fmt.Println(result)
	}
}
