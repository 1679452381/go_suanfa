/**
 * @author X
 * @date 2023/3/22
 * @description
 **/
package main

import "fmt"

func lastStoneWeightII(stones []int) int {

	sum := 0
	for _, stone := range stones {
		sum += stone
	}
	target := sum / 2
	dp := make([]int, target+1)

	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			if dp[j] < dp[j-stones[i]]+stones[i] {
				dp[j] = dp[j-stones[i]] + stones[i]
			}
		}
	}
	return sum - dp[target] - dp[target]
}

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	fmt.Println(lastStoneWeightII(stones))
}
