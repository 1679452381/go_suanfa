package main

import "fmt"

// 由于同一个字母只能出现在同一个片段，显然同一个字母的第一次出现的下标位置和最后一次出现的下标位置必须出现在同一个片段。因此需要遍历字符串，得到每个字母最后一次出现的下标位置。
//
// 在得到每个字母最后一次出现的下标位置之后，可以使用贪心的方法将字符串划分为尽可能多的片段
func partitionLabels(s string) []int {
	endPost := [26]int{}
	res := make([]int, 0)
	for i, v := range s {
		endPost[v-'a'] = i
	}
	fmt.Println(endPost)
	start, end := 0, 0
	for i, v := range s {
		if endPost[v-'a'] > end {
			end = endPost[v-'a']
		}
		if i == end {
			res = append(res, end-start+1)
			start = end + 1
		}
	}
	return res
}

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij")) //"ababcbaca"、"defegde"、"hijhklij" 。
}
