package main

// 对称二叉树
func isSymmetric(root *TreeNode) bool {
	q := make([]*TreeNode, 0)
	q = append(q, root.Left, root.Right)
	for len(q) > 0 {
		left := q[0]
		right := q[1]
		q = q[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		q = append(q, left.Left, right.Right, left.Right, right.Left)
	}
	return false
}

func main() {

}
