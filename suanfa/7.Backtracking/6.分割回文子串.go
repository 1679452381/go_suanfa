/**
 * @author X
 * @date 2023/5/9
 * @description
 **/
package main

import "fmt"

func partition(s string) [][]string {
	res := make([][]string, 0)
	path := make([]string, 0)
	partitionBack(s, &res, path, 0)
	return res
}

func partitionBack(s string, res *[][]string, path []string, start int) {
	if start == len(s) {
		temp := make([]string, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	for i := start; i < len(s); i++ {
		if !isPalindromic(s, start, i) {
			continue
		}
		path = append(path, string([]byte(s)[start:i+1]))
		partitionBack(s, res, path, i+1)
		path = path[:len(path)-1]
	}
}

// 判断是否是回文字符串
func isPalindromic(s string, start, end int) bool {
	for start <= end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func main() {
	s := "aab"
	fmt.Println(partition(s))
}
