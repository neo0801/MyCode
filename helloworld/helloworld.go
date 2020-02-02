package main

import "fmt"

const prefixWord = "Hello, "

// Hello just return hello world!
func Hello(name string) string {
	return prefixWord + name
}

func main() {
	fmt.Println(Hello("world"))
}
