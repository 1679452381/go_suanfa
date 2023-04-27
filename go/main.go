package main

import "fmt"

func main() {
	m := make(map[int]string, 0)
	m[1] = "zxc"
	for k, v := range m {
		fmt.Println(k, v)
	}
}
