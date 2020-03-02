package main

import (
	"bufio"
	"fmt"
	"os"

	fib "github.com/neo0801/MyCode/learn/fibonacci"
)

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	defer fmt.Println("will print latest")
	defer fmt.Println("will print later")
	fmt.Println("will print firstly")
	writeFile("fib.txt")
}
