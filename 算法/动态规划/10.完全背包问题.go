/**
 * @author X
 * @date 2023/3/27
 * @description  完全背包问题
 **/
package main

import "fmt"

func completeBagProble(wight, value []int, bagSize int) int {
	dp := make([]int, bagSize+1)
	for j := wight[0]; j <= bagSize; j++ {
		dp[j] = value[0]
	}
	fmt.Println("dp初始化：", dp)
	for i := 0; i < len(wight); i++ {
		for j := wight[i]; j <= bagSize; j++ {
			if dp[j] < dp[j-wight[i]]+value[i] {
				dp[j] = dp[j-wight[i]] + value[i]
			}
		}
	}
	return dp[bagSize]
}

func main() {
	wight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	bagSize := 4
	maxValue := completeBagProble(wight, value, bagSize)
	fmt.Println(maxValue)
}
