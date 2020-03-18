package main

import "fmt"

func main() {
	fmt.Println(isValid("(){[()]()}()"))
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	stack, size := []rune{}, 0
	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
			size++
		} else {
			if len(stack) == 0 {
				return false
			}
			rear := stack[size-1]
			stack = stack[:size-1]
			switch ch {
			case ')':
				if rear != '(' {
					return false
				}
			case ']':
				if rear != '[' {
					return false
				}
			case '}':
				if rear != '{' {
					return false
				}
			}
			size--
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
