package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return 0
	}
	max := 0
	n := 0
	left := -1
	mapL := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		if i >= 1 {
			delete(mapL, s[i-1])
			n--
		}
		for left+1 < len(s) && mapL[s[left+1]] == 0 {
			mapL[s[left+1]]++
			n++
			left++
		}
		if max < n {
			max = n
		}
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
