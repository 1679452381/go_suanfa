package main

import "fmt"

func bagProblem(wight, value []int, bagSize int) int {
	// 初始化
	dp := make([][]int, len(wight))
	for i, _ := range dp {
		dp[i] = make([]int, bagSize+1)
	}
	for j := wight[0]; j <= bagSize; j++ {
		dp[0][j] = value[0]
	}

	fmt.Println("dp初始化：", dp)
	for i := 1; i < len(wight); i++ {
		for j := 1; j <= bagSize; j++ {
			if j < wight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-wight[i]]+value[i])
			}
		}
	}
	return dp[len(wight)-1][bagSize]
}

// 滚动数据 （一维数组）
func bagProblemOptimized(wight, value []int, bagSize int) int {
	dp := make([]int, bagSize+1)
	for j := wight[0]; j <= bagSize; j++ {
		dp[j] = value[0]
	}
	fmt.Println("dp初始化：", dp)
	for i := 0; i < len(wight); i++ {
		for j := bagSize; j > wight[i]; j-- {
			dp[j] = max(dp[j], dp[j-wight[i]]+value[i])
		}
	}
	return dp[bagSize]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//func main() {
//	wight := []int{1, 3, 4}
//	value := []int{15, 20, 30}
//	bagSize := 4
//	maxValue := bagProblem(wight, value, bagSize)
//	fmt.Println(maxValue)
//	maxValue2 := bagProblemOptimized(wight, value, bagSize)
//	fmt.Println(maxValue2)
//
//}
