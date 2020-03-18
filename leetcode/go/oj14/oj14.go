package main

import "fmt"

func main() {
	strs := []string{"flower", "flow", "aaflight"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	n := len(strs)
	pre := strs[0]
	for i := 1; i < n; i++ {
		cur := strs[i]
		pre = commonPrefix(pre, cur)
	}
	return pre
}

func commonPrefix(s1, s2 string) string {
	n1, n2, n := len(s1), len(s2), 0
	if n1 < n2 {
		n = n1
	} else {
		n = n2
	}
	i := 0
	for ; i < n; i++ {
		if s1[i] != s2[i] {
			break
		}
	}
	return s1[:i]
}
