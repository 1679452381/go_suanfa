package main

import "fmt"

// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
func combine(n int, k int) [][]int {
	ret := make([][]int, 0)
	path := make([]int, 0, k)
	back(&ret, path, n, k, 1)
	return ret

}

func back(ret *[][]int, path []int, n, k, start int) {
	if len(path) == k {
		temp := make([]int, k)
		copy(temp, path)
		*ret = append(*ret, temp)
		return
	}
	for i := start; i <= n-(k-len(path))+1; i++ {
		//for i := start; i <= n; i++ {
		path = append(path, i)
		back(ret, path, n, k, i+1)
		path = path[0 : len(path)-1]
	}
}

func main() {
	ret := combine(4, 2)
	fmt.Println(ret)
}
