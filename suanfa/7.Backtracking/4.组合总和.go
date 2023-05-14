package main

import "fmt"

//给你一个 无重复元素 的整数数组candidates 和一个目标整数target，找出candidates中可以使数字和为目标数target 的 所有不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
//
//candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	combinationSumBack(&res, path, candidates, target, 0)
	return res
}
func combinationSumBack(res *[][]int, path, candidates []int, target, start int) {
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
	}
	for i := start; i < len(candidates); i++ {
		//剪枝优化
		if candidates[i] > target {
			continue
		}
		path = append(path, candidates[i])
		combinationSumBack(res, path, candidates, target-candidates[i], i)
		path = path[:len(path)-1]
	}
}
func main() {
	candidates := []int{8, 7, 4, 3}
	target := 11
	fmt.Println(combinationSum(candidates, target))
}
