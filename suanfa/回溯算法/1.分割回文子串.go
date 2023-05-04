/**
 * @author X
 * @date 2023/4/2
 * @description
 **/
package main

//import "fmt"
//
//var (
//	path []string
//	res  [][]string
//)
//
//func partition(s string) [][]string {
//	path = make([]string, 0)
//	res = make([][]string, 0)
//	dfs(s, 0)
//	return res
//}
//func dfs(s string, startIdx int) {
//	if startIdx == len(s) {
//		tmp := make([]string, len(path))
//		copy(tmp, path)
//		res = append(res, tmp)
//		return
//	}
//
//	for i := startIdx; i < len(s); i++ {
//		str := s[startIdx : i+1]
//		if isPalindrome(str) {
//			path = append(path, str)
//			dfs(s, i+1)
//			path = path[0 : len(path)-1]
//		}
//	}
//}
//
//func isPalindrome(s string) bool {
//	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
//		if s[i] != s[j] {
//			return false
//		}
//	}
//	return true
//}
//
//func main() {
//	s := "aab"
//	fmt.Println(partiton(s))
//
//}
