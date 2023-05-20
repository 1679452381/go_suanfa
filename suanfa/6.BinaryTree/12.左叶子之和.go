package main

import "fmt"

func sumOfLeftLeaves(root *TreeNode) int {
	leftNodes := make([]int, 0)
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur.Left != nil {
			q = append(q, cur.Left)
			if cur.Left.Left == nil && cur.Left.Right == nil {
				leftNodes = append(leftNodes, cur.Left.Val)
			}
		}
		if cur.Right != nil {
			q = append(q, cur.Right)
		}
	}
	fmt.Println(leftNodes)
	sum := 0
	for _, node := range leftNodes {
		sum += node
	}
	return sum
}
