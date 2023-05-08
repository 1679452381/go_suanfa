package main

import (
	"fmt"
	"sort"
)

//给定一个候选人编号的集合 candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。
//
//candidates中的每个数字在每个组合中只能使用一次。

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	if len(candidates) == 0 {
		return res
	}
	used := make([]bool, len(candidates))
	sort.Ints(candidates)
	combinationSum2Back(&res, path, candidates, target, 0, used)
	return res
}

func combinationSum2Back(res *[][]int, path, candidates []int, target, start int, used []bool) {
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
	}
	for i := start; i < len(candidates); i++ {
		if i > 0 && candidates[i] == candidates[i-1] && used[i-1] != true {
			continue
		}
		//剪枝优化
		if candidates[i] > target {
			return
		}
		used[i] = true
		path = append(path, candidates[i])
		combinationSum2Back(res, path, candidates, target-candidates[i], i+1, used)
		path = path[:len(path)-1]
		used[i] = false
	}
}
func main() {
	candidates := []int{1, 1, 2, 5, 6, 7, 10}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}
