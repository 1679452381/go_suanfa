package main

import "fmt"

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(nums))
	permuteBack(nums, path, used, len(nums), &res)
	return res
}

func permuteBack(nums, path []int, used []bool, numsl int, res *[][]int) {
	if len(path) == numsl {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	for i := 0; i < numsl; i++ {
		if used[i] {
			continue
		}
		used[i] = true
		path = append(path, nums[i])
		permuteBack(nums, path, used, numsl, res)
		path = path[0 : len(path)-1]
		used[i] = false
	}
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
