/**
 * @author X
 * @date 2023/3/22
 * @description
 **/
package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}
	fmt.Println(dp)

	for i := 0; i < m && obstacleGrid[i][0] != 1; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n && obstacleGrid[0][j] != 1; j++ {
		dp[0][j] = 1
	}
	fmt.Println(dp)
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	fmt.Println(dp)
	return dp[m-1][n-1]
}

func main() {
	//obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	obstacleGrid := [][]int{{0, 0}}
	res := uniquePathsWithObstacles(obstacleGrid)
	fmt.Println(res)
}
