package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(-123))
}

func reverse(x int) int {
	var tmp int
	if x < 0 {
		tmp = -1 * reversePostive(-1*x)
		if tmp < -1*int(math.Pow(2, 31)) {
			tmp = 0
		}
	} else {
		tmp = reversePostive(x)
		if tmp >= int(math.Pow(2, 31)) {
			tmp = 0
		}
	}
	return tmp
}

func reversePostive(x int) int {
	res := 0
	for x > 0 {
		res *= 10
		res += x % 10
		x /= 10
	}
	return res
}
