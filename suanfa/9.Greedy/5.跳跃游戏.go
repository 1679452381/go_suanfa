package main

import (
	"fmt"
)

func canJump(nums []int) bool {
	cover := 0
	for i := 0; i <= cover; i++ {
		if cover < nums[i]+i {
			cover = nums[i] + i
		}
		if cover >= len(nums)-1 {
			return true
		}
	}
	return false
}
func main() {
	nums := []int{3, 2, 1, 1, 4}
	fmt.Println(canJump(nums))
}
