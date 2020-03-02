package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fibonacci() fibGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type fibGen func() int

func (f fibGen) Read(p []byte) (n int, err error) {
	next := f()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printIObuffer(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()
	printIObuffer(f)
}
