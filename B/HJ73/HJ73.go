package main

import (
	"fmt"
	"time"
)

/*
题目描述
根据输入的日期，计算是这一年的第几天。。

详细描述：

输入某年某月某日，判断这一天是这一年的第几天？。

测试用例有多组，注意循环输入

接口设计及说明：
*****************************************************************************
Description   : 数据转换
Input Param   : year 输入年份
Month 输入月份
Day 输入天

Output Param  :
Return Value  : 成功返回0，失败返回-1（如：数据错误）
*****************************************************************************
***************************************************************************
  Description   :
  Input Param   :

  Output Param  :
  Return Value  : 成功:返回outDay输出计算后的第几天;
  失败:返回-1
***************************************************************************
输入描述:
输入多行，每行空格分割，分别是年，月，日

输出描述:
成功:返回outDay输出计算后的第几天;
失败:返回-1

示例1
输入
复制
2012 12 31
输出
复制
366

*/
func main() {
	for {
		y, m, d := 0, 0, 0
		_, err := fmt.Scan(&y, &m, &d)
		if err != nil {
			break
		}
		t, err := time.Parse("2006-01-02", fmt.Sprintf("%04d-%02d-%02d", y, m, d))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(t.YearDay())
	}
}
