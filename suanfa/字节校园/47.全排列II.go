/**
 * @author X
 * @date 2023/6/26
 * @description
 **/
package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(nums))
	permuteUniqueBack(nums, path, used, len(nums), &res)
	return res
}

func permuteUniqueBack(nums, path []int, used []bool, numsl int, res *[][]int) {
	if len(path) == numsl {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	for i := 0; i < numsl; i++ {
		if used[i] || i > 0 && nums[i] == nums[i-1] && used[i-1] != true {
			continue
		}
		used[i] = true
		path = append(path, nums[i])
		permuteUniqueBack(nums, path, used, numsl, res)
		path = path[0 : len(path)-1]
		used[i] = false
	}
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 3}))
}
