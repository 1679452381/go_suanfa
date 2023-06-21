/**
 * @author X
 * @date 2023/6/15
 * @description
 **/
package main

import "fmt"

func bubbleSort(nums []int) []int {
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		for j := 0; j < numsLen-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func main() {
	nums := bubbleSort([]int{2, 3, 1, 5, 56, 82, 6, 9})
	fmt.Println(nums)
}
