package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	record := [26]int{}
	for _, val := range s {
		record[val-rune('a')]++
	}
	for _, val := range t {
		record[val-rune('a')]--
	}
	return record == [26]int{}
}

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
}
