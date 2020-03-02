package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%#v\n", strings.Fields("  s string   a     db"))
	fmt.Printf("%#v\n", strings.Split("  s string   a     db", ""))
	fmt.Println(lenOfNonRepeatSubString(""))
	fmt.Println(lenOfNonRepeatSubString("eng"))
	fmt.Println(lenOfNonRepeatSubString("中文测试"))
	fmt.Println(lenOfNonRepeatSubString("中文测试中"))
	fmt.Println(lenOfNonRepeatSubString("中文测试中，不可中断"))
	fmt.Println(lenOfNonRepeatSubString("黑化肥挥发发黑"))
}

func lenOfNonRepeatSubString(s string) int {
	lastOccured := make(map[rune]int)
	start, maxLength := 0, 0
	for i, ch := range []rune(s) {
		if lo, ok := lastOccured[ch]; ok && lo >= start {
			start = lo + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccured[ch] = i
	}
	return maxLength
}
