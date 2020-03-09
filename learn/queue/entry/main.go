package main

import (
	"fmt"

	queue "github.com/neo0801/MyCode/learn/queue"
	"golang.org/x/tools/container/intsets"
)

func testSparse() {
	s := intsets.Sparse{}
	s.Insert(1)
	s.Insert(1000)
	s.Insert(1000000000)
	fmt.Println(s.Has(1000))
	fmt.Println(s.Has(10000000))
}

func main() {
	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	testSparse()
}
