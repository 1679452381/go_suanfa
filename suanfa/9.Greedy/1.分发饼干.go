package main

import (
	"fmt"
	"sort"
)

// 分发饼干
func findContentChildren(g []int, s []int) int {
	count := 0
	sort.Ints(g)
	sort.Ints(s)
	lenS := len(s) - 1
	lenG := len(g) - 1
	for lenS >= 0 && lenG >= 0 {
		if s[lenS] >= g[lenG] {
			count++
			lenS--
		}
		//g = g[:lenG]
		lenG--
	}
	return count
}

func main() {
	g := []int{1, 2, 3}
	s := []int{1, 1}
	res := findContentChildren(g, s)
	fmt.Println(res)

}
