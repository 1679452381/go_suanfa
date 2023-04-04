/**
 * @author X
 * @date 2023/3/24
 * @description
 **/
package main

import "fmt"

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	for _, str := range strs {
		zeroNum := 0
		oneNum := 0
		for _, v := range str {
			if v == '0' {
				zeroNum++
			} else {
				oneNum++
			}
		}
		//fmt.Println(zeroNum, oneNum)
		for i := m; i >= zeroNum; i-- {
			for j := n; j >= oneNum; j-- {
				if dp[i][j] < dp[i-zeroNum][j-oneNum]+1 {
					dp[i][j] = dp[i-zeroNum][j-oneNum] + 1
				}
			}
		}
		//fmt.Println(dp)
	}
	return dp[m][n]
}

func main() {
	strs := []string{"10", "0001", "111001", "1", "0"}
	m := 3
	n := 3
	fmt.Println(findMaxForm(strs, m, n))
}
