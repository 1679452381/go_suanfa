package main

import "fmt"

// 使用中序遍历判断是否为二叉搜索树

func isValidBST(root *TreeNode) bool {
	s := make([]*TreeNode, 0)
	cur := root
	var pre *TreeNode
	for len(s) > 0 || cur != nil {
		for cur != nil {
			s = append(s, cur)
			cur = cur.Left
		}
		cur = s[len(s)-1]
		s = s[:len(s)-1]
		if pre != nil && cur.Val <= pre.Val {
			return false
		}
		pre = cur
		cur = cur.Right
	}
	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 0}
	fmt.Println(isValidBST(root))
}
