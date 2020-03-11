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

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d recieved %d.\n", id, n)
	}
}

func createWorker(id int) chan int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	c1, c2, worker := generator(), generator(), createWorker(0)
	values := []int{}
	tm := time.After(time.Minute)
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-tm:
			fmt.Println("Bye!")
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Time out.")
		case <-tick:
			fmt.Println("Array len:", len(values))
		}
	}
}
