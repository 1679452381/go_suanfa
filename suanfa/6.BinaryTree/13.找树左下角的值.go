package main

func findBottomLeftValue(root *TreeNode) int {
	res := root.Val
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			cur := q[0]
			q = q[1:]
			if i == 0 {
				res = cur.Val
			}
			if cur.Left != nil {
				q = append(q, cur.Left)
				if cur.Left.Left == nil && cur.Left.Right == nil {
					res = cur.Left.Val
				}
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
	}
	return res
}
