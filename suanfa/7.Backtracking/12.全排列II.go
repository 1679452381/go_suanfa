package main

import "fmt"

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	numsLen := len(nums)
	path := make([]int, 0, numsLen)
	used := make(map[int]bool)
	permuteUniqueBack(&res, path, nums, numsLen, used)
	return res
}
func permuteUniqueBack(res *[][]int, path, nums []int, numsLen int, used map[int]bool) {
	if len(path) == numsLen {
		temp := make([]int, numsLen)
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	for i := 0; i < numsLen; i++ {
		if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i]) {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		permuteUniqueBack(res, path, nums, numsLen, used)
		used[i] = false
		path = path[:len(path)-1]
	}

}

func main() {
	nums := []int{1, 1, 2}
	res := permuteUnique(nums)
	fmt.Println(res)
}
