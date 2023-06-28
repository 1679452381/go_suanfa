/**
 * @author X
 * @date 2023/6/25
 * @description
 **/
package main

import "fmt"

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	res := s[0:1]
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] == true && j-i+1 > len(res) {
				res = s[i : j+1]
			}
		}
	}
	return res
}
func main() {
	fmt.Println(longestPalindrome("aaaaaa"))
}

//
//[false true true false false false]
//[false true true true false false]
//[false false true true true false]
//[false false false true true true]
//[false false false false true true]
//[false false false false false true]
//
//[[true false false false false false]
//[false true false false false false]
//[false false true false false false]
//[false false false true false false]
//[false false false false true false]
//[false false false false false true]]
