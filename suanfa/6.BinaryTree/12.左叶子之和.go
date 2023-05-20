package main

func sumOfLeftLeaves(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	sum := 0
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur.Left != nil {
			q = append(q, cur.Left)
			if cur.Left.Left == nil && cur.Left.Right == nil {
				sum += cur.Left.Val
			}
		}
		if cur.Right != nil {
			q = append(q, cur.Right)
		}
	}
	return sum
}
