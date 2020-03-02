package main

import (
	"fmt"

	retriever "github.com/neo0801/MyCode/learn/interfeces/Retrirver"
)

// Retriever is
type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("url string)")
}

func main() {
	r := retriever.Retriever{"fake news!"}
	fmt.Println(r.Get("华尔街日报"))
}
