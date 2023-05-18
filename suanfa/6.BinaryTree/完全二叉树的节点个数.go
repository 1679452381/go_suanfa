package main

import (
	"fmt"
)

func countNodesRe(root *TreeNode) int {
	count := 0
	if root == nil {
		return count
	}
	left := root.Left
	right := root.Right
	leftDep := 0
	for left != nil {
		leftDep++
		left = left.Left
	}
	rightDep := 0
	for right != nil {
		rightDep++
		right = right.Right
	}
	if leftDep == rightDep {
		return (2 << leftDep) - 1
	}
	return countNodesRe(root.Left) + countNodes(root.Right) + 1
}

func countNodes(root *TreeNode) int {
	count := 0
	if root == nil {
		return count
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			cur := q[0]
			count++
			q = q[1:]
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
	}
	return count
}

func main() {
	fmt.Println(2 ^ 4)
}
