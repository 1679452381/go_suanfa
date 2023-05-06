package main

import "fmt"

func reverseString(s []byte) string {
	start := 0
	end := len(s) - 1
	for start <= end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
	return string(s)
}
func main() {
	s := []byte("123456")
	fmt.Println(reverseString(s))
}
