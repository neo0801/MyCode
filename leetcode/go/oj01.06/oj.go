package main

import (
	"fmt"
)

func main() {
	fmt.Println(compressString("rrrrrLLLLLPPPPPPRRRRRgggNNNNNVVVVVVVVVVDDDDDDDDDDIIIIIIIIIIlllllllAAAAqqqqqqqbbbNNNNffffff"))
	fmt.Println(compressString("abbccd"))
}

func compressString(S string) string {
	res := compress(S)
	if len(res) < len(S) {
		return res
	}
	return S
}

func compress(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	pre := s[0]
	res, cnt := "", 1
	res = fmt.Sprintf("%s%s", res, pre)
	for i := 1; i < n; i++ {
		cur := s[i]
		if cur != pre {
			res = fmt.Sprintf("%s%d", res, cnt)
			pre = cur
			res = fmt.Sprintf("%s%s", res, cur)
			cnt = 1
		} else {
			cnt++
		}
	}
	res = fmt.Sprintf("%s%d", res, cnt)
	return res
}
