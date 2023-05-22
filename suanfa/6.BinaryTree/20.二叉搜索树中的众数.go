package main

import "fmt"

func findMode(root *TreeNode) []int {
	s := make([]*TreeNode, 0)
	res := make([]int, 0)
	cur := root
	count := 0
	maxCount := 0
	var pre *TreeNode
	for len(s) > 0 || cur != nil {
		for cur != nil {
			s = append(s, cur)
			cur = cur.Left
		}
		cur = s[len(s)-1]
		s = s[:len(s)-1]

		if pre == nil {
			count = 1
		} else if pre.Val == cur.Val {
			count++
		} else {
			count = 1
		}
		if count == maxCount {
			res = append(res, cur.Val)
		}
		if count > maxCount {
			maxCount = count
			res = res[:0]
			res = append(res, cur.Val)
		}
		pre = cur
		cur = cur.Right
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1,
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 2},
			Right: nil,
		}}
	root2 := &TreeNode{Val: 0}
	fmt.Println(findMode(root))
	fmt.Println(findMode(root2))
}
