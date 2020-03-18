package main

import (
	"fmt"
)

func main() {
	rec1 := []int{0, 0, 2, 2}
	rec2 := []int{1, 1, 3, 3}
	fmt.Println(isRectangleOverlap(rec1, rec2))
	rec1 = []int{0, 0, 1, 1}
	rec2 = []int{1, 0, 2, 1}
	fmt.Println(isRectangleOverlap(rec1, rec2))
}

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// [x1, y1, x2, y2] rec1
	// [x3, y3, x4, y4] rec2
	// !((x2<=x3 || x4<=x1) || (y2<=y3 || y4<=y1))
	return !(rec1[2] <= rec2[0] || rec2[2]<= rec1[0] || rec1[3] <= rec2[1] || rec2[3] <= rec1[1])
}
