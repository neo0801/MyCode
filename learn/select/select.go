package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func main() {
	c1, c2, n := generator(), generator(), 0
	for {
		select {
		case n = <-c1:
			fmt.Println("Recieved from channel c1:", n)
		case n = <-c2:
			fmt.Println("Recieved from channel c2:", n)
		}
		if n > 10 {
			fmt.Println("Done!")
			break
		} else {
			fmt.Println("continue with n:", n)
		}
	}
}
