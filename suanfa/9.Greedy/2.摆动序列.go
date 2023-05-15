/**
 * @author X
 * @date 2023/5/15
 * @description
 **/
package main

import "fmt"

func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	preDiff := 0
	count := 1
	for i := 0; i < len(nums)-1; i++ {
		diff := nums[i+1] - nums[i]
		if diff > 0 && preDiff <= 0 || diff < 0 && preDiff >= 0 {
			preDiff = diff
			count++
		}
	}
	return count
}
func main() {
	nums := []int{1, 7, 4, 9, 2, 5}
	res := wiggleMaxLength(nums)
	fmt.Println(res)

}
