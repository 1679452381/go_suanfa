package main

import (
	"fmt"
)

var numberString []string = []string{
	"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
}

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	path := make([]byte, 0, len(digits))
	if len(digits) == 0 {
		return res
	}
	letterCombinationsBack(&res, digits, path, 0)
	return res
}
func letterCombinationsBack(res *[]string, digits string, path []byte, start int) {
	if len(path) == len(digits) {
		temp := string(path)
		*res = append(*res, temp)
		return
	}
	digit := int(digits[start] - '0')
	str := numberString[digit-2]
	for i := 0; i < len(str); i++ {
		path = append(path, str[i])
		letterCombinationsBack(res, digits, path, start+1)
		path = path[:len(path)-1]
	}

}

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}
