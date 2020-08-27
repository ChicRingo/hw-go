package main

import "fmt"

/*
题目描述
如果A是个x行y列的矩阵，B是个y行z列的矩阵，把A和B相乘，其结果将是另一个x行z列的矩阵C。这个矩阵的每个元素是由下面的公式决定的



输入描述:
输入包含多组数据，每组数据包含：

第一行包含一个正整数x，代表第一个矩阵的行数

第二行包含一个正整数y，代表第一个矩阵的列数和第二个矩阵的行数

第三行包含一个正整数z，代表第二个矩阵的列数

之后x行，每行y个整数，代表第一个矩阵的值

之后y行，每行z个整数，代表第二个矩阵的值



输出描述:
对于每组输入数据，输出x行，每行z个整数，代表两个矩阵相乘的结果
示例1
输入
复制
2
3
2
1 2 3
3 2 1
1 2
2 1
3 3
输出
复制
14 13
10 11
*/
func main() {
	for {
		r, rc, c := 0, 0, 0
		if _, err := fmt.Scan(&r, &rc, &c); err != nil {
			return
		}

		arr1 := make([][]int, r)
		for i := 0; i < r; i++ {
			arr1[i] = make([]int, rc)
		}
		arr2 := make([][]int, rc)
		for i := 0; i < rc; i++ {
			arr2[i] = make([]int, c)
		}
		for i := 0; i < r; i++ {
			for j := 0; j < rc; j++ {
				fmt.Scan(&arr1[i][j])
			}
		}
		for i := 0; i < rc; i++ {
			for j := 0; j < c; j++ {
				fmt.Scan(&arr2[i][j])
			}
		}

		res := make([][]int, r)
		for i := 0; i < r; i++ {
			res[i] = make([]int, c)
		}

		for i := 0; i < r; i++ {
			for j := 0; j < c; j++ {
				for k := 0; k < rc; k++ {
					res[i][j] += arr1[i][k] * arr2[k][j]
				}
			}
		}

		for i := 0; i < r; i++ {
			for j := 0; j < c; j++ {
				fmt.Printf("%d ", res[i][j])
			}
			fmt.Println()
		}
	}
}
