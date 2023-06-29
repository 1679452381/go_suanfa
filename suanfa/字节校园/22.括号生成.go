package main

import "fmt"

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	path := make([]byte, 0)
	generatePBack(&res, path, 0, 0, n)
	return res
}

func generatePBack(res *[]string, path []byte, left, right, n int) {
	if len(path) == n*2 {
		temp := make([]byte, len(path))
		copy(temp, path)
		*res = append(*res, string(temp))
	}
	if left < n {
		path = append(path, '(')
		generatePBack(res, path, left+1, right, n)
		path = path[:len(path)-1]
	}
	if right < left {
		path = append(path, ')')
		generatePBack(res, path, left, right+1, n)
		path = path[:len(path)-1]
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}
