package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()

	return a.value
}

func (a *atomicInt) increment() {
	a.lock.Lock()
	defer a.lock.Unlock()

	a.value++
}

func main() {
	var a atomicInt
	a.increment()
	for i := 0; i < 10000; i++ {
		go func() {
			a.increment()
		}()
	}
	time.Sleep(time.Millisecond * 200)
	fmt.Println(a.get())
}
