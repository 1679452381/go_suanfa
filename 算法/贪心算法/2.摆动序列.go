/**
 * @author X
 * @date 2023/3/22
 * @description
 **/
package main

import "fmt"

func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	res := 1
	prediff := 0
	curdiff := 0
	for i := 0; i < len(nums)-1; i++ {
		curdiff = nums[i+1] - nums[i]
		if (prediff <= 0 && curdiff > 0) || (prediff >= 0 && curdiff < 0) {
			res++
			prediff = curdiff
		}
	}
	return res
}

func main() {
	nums := []int{1, 7, 4, 9, 2, 5}
	fmt.Println(wiggleMaxLength(nums))
}
