package main

// 二叉树的最大深度
// 层序遍历
func maxDepth(root *TreeNode) int {
	max := 0
	if root == nil {
		return max
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		max++
		n := len(q)
		for i := 0; i < n; i++ {
			cur := q[0]
			q = q[1:]
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
	}
	return max
}

func main() {

}
