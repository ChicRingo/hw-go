package main

import "fmt"

/*
题目描述
定义一个二维数组N*M（其中2<=N<=10;2<=M<=10），如5 × 5数组下所示：


int maze[5][5] = {


        0, 1, 0, 0, 0,


        0, 1, 0, 1, 0,


        0, 0, 0, 0, 0,


        0, 1, 1, 1, 0,


        0, 0, 0, 1, 0,


};


它表示一个迷宫，其中的1表示墙壁，0表示可以走的路，只能横着走或竖着走，不能斜着走，要求编程序找出从左上角到右下角的最短路线。入口点为[0,0],既第一空格是可以走的路。

Input

一个N × M的二维数组，表示一个迷宫。数据保证有唯一解,不考虑有多解的情况，即迷宫只有一条通道。

Output

左上角到右下角的最短路径，格式如样例所示。

Sample Input

0 1 0 0 0

0 1 0 1 0

0 0 0 0 0

0 1 1 1 0

0 0 0 1 0

Sample Output

(0, 0)

(1, 0)

(2, 0)

(2, 1)

(2, 2)

(2, 3)

(2, 4)

(3, 4)

(4, 4)

输入描述:
输入两个整数，分别表示二位数组的行数，列数。再输入相应的数组，其中的1表示墙壁，0表示可以走的路。数据保证有唯一解,不考虑有多解的情况，即迷宫只有一条通道。

输出描述:
左上角到右下角的最短路径，格式如样例所示。

示例1
输入
复制
5 5
0 1 0 0 0
0 1 0 1 0
0 0 0 0 0
0 1 1 1 0
0 0 0 1 0
输出
复制
(0,0)
(1,0)
(2,0)
(2,1)
(2,2)
(2,3)
(2,4)
(3,4)
(4,4)
*/
type P struct {
	i, j int
}

var (
	N, M int
	m    [][]int
	best []P
	used []P
)

func main() {
	for {
		_, err := fmt.Scan(&N, &M)
		if err != nil || N < 1 || M < 1 {
			break
		}
		best = []P{}
		used = []P{}
		m = make([][]int, N)
		for i := 0; i < N; i++ {
			m[i] = make([]int, M)
			for j := 0; j < M; j++ {
				fmt.Scan(&m[i][j])
			}
		}
		//fmt.Println(N,M,m)
		Track(0, 0)
		for _, p := range best {
			fmt.Printf("(%d,%d)\n", p.i, p.j)
		}
	}
}
func Track(i, j int) {
	m[i][j] = 1
	used = append(used, P{i, j})
	if i == N-1 && j == M-1 {
		if len(best) == 0 || len(used) < len(best) {
			best = used
		}
		used = used[:len(used)-1]
		m[i][j] = 0
		return
	}

	if i-1 >= 0 && m[i-1][j] == 0 { //向上可走
		Track(i-1, j)
	}
	if i+1 < N && m[i+1][j] == 0 { //向下可走
		Track(i+1, j)
	}
	if j-1 >= 0 && m[i][j-1] == 0 { //向左可走
		Track(i, j-1)
	}
	if j+1 < M && m[i][j+1] == 0 { //向右可走
		Track(i, j+1)
	}
	used = used[:len(used)-1]
	m[i][j] = 0
}
