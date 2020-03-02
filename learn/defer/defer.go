package main

import "fmt"

func main() {
	defer fmt.Println("will print latest")
	defer fmt.Println("will print later")
	fmt.Println("will print firstly")
}
