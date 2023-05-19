package main

func isBalanced(root *TreeNode) bool {
	res := getH(root)
	if res == -1 {
		return false
	}
	return true
}
func getH(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := getH(root.Left), getH(root.Right)
	if l == -1 || r == -1 {
		return -1
	}
	if l-r > 1 || r-l > 1 {
		return -1
	}
	if l > r {
		return l + 1
	}
	return r + 1
}
func main() {

}
