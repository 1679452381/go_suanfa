package main

func minDepth(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		n := len(q)
		depth++
		for i := 0; i < n; i++ {
			cur := q[0]
			q = q[1:]
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
		}
	}
	return depth
}
func main() {

}
