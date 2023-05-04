/**
 * @author X
 * @date 2023/4/4
 * @description  多重背包  背包容量 bag_weight  物品价值v[],物品数量n[],物品重量w[] 计算背包能装物品的最大价值
 **/
package main

import "fmt"

// 转化为01背包问题
func MultiPack(bagWeight int, value, nums, weight []int) int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < nums[i]-1; j++ {
			weight = append(weight, weight[i])
			value = append(value, value[i])
		}
	}
	fmt.Println(weight, value)
	dp := make([]int, bagWeight+1)
	for i := 0; i < len(weight); i++ {
		for j := bagWeight; j >= weight[i]; j-- {
			if dp[j] < dp[j-weight[i]]+value[i] {
				dp[j] = dp[j-weight[i]] + value[i]
			}
		}
	}
	return dp[bagWeight]
}

func main() {
	bag_weight := 4
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	nums := []int{2, 3, 2}

	fmt.Println(MultiPack(bag_weight, value, nums, weight))
}
