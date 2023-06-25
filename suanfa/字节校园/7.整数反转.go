package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	res := 0
	for x != 0 {
		temp := x % 10
		if res > math.MaxInt32/10 || res < math.MinInt32/10 {
			return 0
		}

		res = res*10 + temp
		x = x / 10
	}

	return res
}

func main() {
	fmt.Println(reverse(-321))
}
