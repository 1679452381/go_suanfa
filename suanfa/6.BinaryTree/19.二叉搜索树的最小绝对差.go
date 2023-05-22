package main

import (
	"fmt"
	"math"
)

func getMinimumDifference(root *TreeNode) int {
	s := make([]*TreeNode, 0)
	cur := root
	min := math.MaxInt
	var pre *TreeNode
	for len(s) > 0 || cur != nil {
		for cur != nil {
			s = append(s, cur)
			cur = cur.Left
		}
		cur = s[len(s)-1]
		s = s[:len(s)-1]
		if pre != nil {
			if min > cur.Val-pre.Val {
				min = cur.Val - pre.Val
			}
		}
		pre = cur
		cur = cur.Right
	}
	return min
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1,
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 2},
			Right: nil,
		}}
	fmt.Println(getMinimumDifference(root))
}
