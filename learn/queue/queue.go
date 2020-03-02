package queue

// Queue is []int
type Queue []int

// Push a element into Queue
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// Pop first element in Queue
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// IsEmpty returns whether Queue is empty
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
