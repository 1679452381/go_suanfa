package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	n := len(s)
	left := -1
	mapS := make(map[uint8]int)
	for i := 0; i < n; i++ {
		if i > 0 {
			delete(mapS, s[i-1])
		}
		for left+1 < n && mapS[s[left+1]] == 0 {
			mapS[s[left+1]] = 1
			left++
		}
		if maxLen < len(mapS) {
			maxLen = len(mapS)
		}
	}
	return maxLen

}
func main() {
	fmt.Println(lengthOfLongestSubstring("zxczxc"))
}
