package main

import "fmt"

/*
题目描述
问题描述：给出4个1-10的数字，通过加减乘除，得到数字为24就算胜利
输入：
4个1-10的数字。[数字允许重复，但每个数字仅允许使用一次，测试用例保证无异常数字]
输出：
true or false

输入描述:
输入4个int整数

输出描述:
返回能否得到24点，能输出true，不能输出false

示例1
输入
复制
7 2 1 10
输出
复制
true
*/
func main() {
	for {
		var nums []int
		var num int
		for i := 0; i < 4; i++ {
			n, err := fmt.Scan(&num)
			if n < 1 && err != nil {
				return
			}
			nums = append(nums, num)
		}
		fmt.Println(find24(nums, 24))
	}
}

func find24(nums []int, total int) bool {
	if total == 0 {
		return true
	}
	if len(nums) == 1 {
		if nums[0] == total {
			return true
		} else {
			return false
		}
	}

	for i, _ := range nums {
		curV := nums[i]
		leftNums := make([]int, len(nums))
		copy(leftNums, nums)
		leftNums = append(leftNums[:i], leftNums[i+1:]...)

		var b1, b2, b3, b4, b5 bool
		/* + */
		if total >= curV {
			b1 = find24(leftNums, total-curV)
		}
		/* - */
		b2 = find24(leftNums, total+curV)

		/* * */
		if total%curV == 0 {
			b3 = find24(leftNums, total/curV)
		}

		/* / */
		b4 = find24(leftNums, total*curV)

		/* 无操作 */
		b5 = find24(leftNums, total)

		if b1 || b2 || b3 || b4 || b5 {
			return true
		}
	}

	return false
}
