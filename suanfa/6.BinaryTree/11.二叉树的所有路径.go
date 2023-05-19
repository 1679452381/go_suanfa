package main

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	pathQ := make([]string, 0)
	nodeQ := make([]*TreeNode, 0)
	nodeQ = append(nodeQ, root)
	pathQ = append(pathQ, strconv.Itoa(root.Val))
	for len(nodeQ) > 0 {
		cur := nodeQ[0]
		nodeQ = nodeQ[1:]
		path := pathQ[0]
		pathQ = pathQ[1:]
		if cur.Left == nil && cur.Right == nil {
			res = append(res, path)
		}
		if cur.Left != nil {
			nodeQ = append(nodeQ, cur.Left)
			pathQ = append(pathQ, path+"->"+strconv.Itoa(cur.Left.Val))
		}
		if cur.Right != nil {
			nodeQ = append(nodeQ, cur.Right)
			pathQ = append(pathQ, path+"->"+strconv.Itoa(cur.Right.Val))
		}
	}
	return res
}
func main() {

}
