/**
 * @author X
 * @date 2023/5/9
 * @description
 **/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	path := make([]string, 0)
	restoreIpAddressesBack(s, &res, path, 0)
	return res
}
func restoreIpAddressesBack(s string, res *[]string, path []string, start int) {
	if len(path) == 4 {
		if start == len(s) {
			temp := strings.Join(path, ".")
			//fmt.Println(temp)
			*res = append(*res, temp)
		}
		return
	}
	for i := start; i < len(s); i++ {
		if i != start && s[start] == '0' {
			break
		}
		str := s[start : i+1]
		num, _ := strconv.Atoi(str)
		if num >= 0 && num <= 255 {
			path = append(path, str)
			restoreIpAddressesBack(s, res, path, i+1)
			path = path[:len(path)-1]
		} else {
			break
		}
	}
}

func main() {
	s := "101023"
	fmt.Println(restoreIpAddresses(s))
}
