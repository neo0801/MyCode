package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

var dict = map[string]int{
	"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
}

func romanToInt(s string) int {
	n, m := len(s), 0
	for i := n - 1; i >= 0; i-- {
		value := dict[string(s[i])]
		switch {
		case m >= 1000:
			if value == 1000 {
				m += value
			} else {
				m -= value
			}
		case m >= 500:
			if value >= 500 {
				m += value
			} else {
				m -= value
			}
		case m >= 100:
			if value >= 100 {
				m += value
			} else {
				m -= value
			}
		case m >= 50:
			if value >= 50 {
				m += value
			} else {
				m -= value
			}
		case m >= 10:
			if value >= 10 {
				m += value
			} else {
				m -= value
			}
		case m >= 5:
			if value >= 5 {
				m += value
			} else {
				m -= value
			}
		case m >= 0:
			m += value
		}
	}
	return m
}
