/**
 * @author X
 * @date 2023/5/15
 * @description
 **/
package main

import "fmt"

func maxSubArray(nums []int) int {
	res := 0
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > res+nums[i] {
			res = nums[i]
		} else {
			res = res + nums[i]
		}
		if res > max {
			max = res
		}
	}
	return max
}

func main() {
	//nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	nums := []int{-1}
	fmt.Println(maxSubArray(nums))
}
