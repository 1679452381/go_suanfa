package main

import "fmt"

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == q || root == p || root == nil {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil && right != nil {
		return right
	} else if left != nil && right == nil {
		return left
	}
	return nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 4},
			},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 8},
		}}
	p := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 6,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: 4},
		},
	}
	q := &TreeNode{Val: 4}
	fmt.Println(lowestCommonAncestor(root, p, q))

}
