package main

import "fmt"

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	numsLen := len(nums)
	path := make([]int, 0, numsLen)
	used := make(map[int]bool)
	permuteBack(&res, path, nums, numsLen, used)
	return res

}
func permuteBack(res *[][]int, path, nums []int, numsLen int, used map[int]bool) {
	if len(path) == numsLen {
		temp := make([]int, numsLen)
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	for i := 0; i < numsLen; i++ {
		if used[nums[i]] {
			continue
		}
		path = append(path, nums[i])
		used[nums[i]] = true
		permuteBack(res, path, nums, numsLen, used)
		used[nums[i]] = false
		path = path[:len(path)-1]
	}

}

func main() {
	nums := []int{1, 2, 3}
	res := permute(nums)
	fmt.Println(res)
}
