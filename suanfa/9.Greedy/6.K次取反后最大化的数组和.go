package main

import (
	"fmt"
	"math"
	"sort"
)

// 按照绝对值从小到大排序
// 将负数取反 k--
// 负数全部取反 k仍然不为0 则将最小的数取反
func largestSumAfterKNegations(nums []int, k int) int {
	maxSum := 0
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		if k != 0 && nums[i] < 0 {
			nums[i] = -nums[i]
			k--
		}
	}
	fmt.Println(nums)
	if k != 0 && k%2 != 0 {
		nums[len(nums)-1] = -nums[len(nums)-1]
	}
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		maxSum += nums[i]
	}
	return maxSum
}

func main() {
	nums := []int{-2, 9, 9, 8, 4}
	fmt.Println(largestSumAfterKNegations(nums, 5))
}
