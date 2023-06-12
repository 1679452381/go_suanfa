/**
 * @author X
 * @date 2023/6/12
 * @description
 **/
package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	path := make([]int, 0, 3)
	res := make([][]int, 0)
	sum := 0
	used := make([]int, len(nums))
	fmt.Println(nums)
	fmt.Println(used)
	threeSumBack(nums, path, used, &res, sum, 0)
	return res
}
func threeSumBack(nums, path, used []int, res *[][]int, sum, start int) {
	if len(path) == 3 {
		if sum == 0 {
			temp := make([]int, 3)
			copy(temp, path)
			*res = append(*res, temp)
		}
		return
	}
	for i := start; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && used[i-1] == 0 {
			continue
		}
		used[i] = 1
		sum += nums[i]
		path = append(path, nums[i])
		threeSumBack(nums, path, used, res, sum, i+1)
		used[i] = 0
		sum -= nums[i]
		path = path[:len(path)-1]
	}
}

func main() {
	nums := []int{0, 0, 0}

	fmt.Println(threeSum(nums))
}
