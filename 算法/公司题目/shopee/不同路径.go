package main

import "fmt"

/**
* @Author : awen
* @Date : 2020/2/28 11:14 下午
 */

func main() {
	res := uniquePaths(3, 2)
	fmt.Println(res)
}

func uniquePaths(m,n int) int{
	var res [][]int

	res[0][0] = 1
	for i := 0; i<m; i ++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				res[i][j] = res[i][j-1]
			} else if j==0 {
				res[i][j] = res[i - 1][j]
			} else {
				res[i][j] = res[i - 1][j] + res[i][j - 1];
			}
		}
	}
	// 到达终点的路径可能性
	return res[m-1][n-1]
}
