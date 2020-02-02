package main

import (
	"fmt"
	"testing"
)

// TestHello tests Hello function
func TestHello(t *testing.T) {
	got := Hello("world!")
	want := "Hello, world!"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i<b.N; i++ {
		fmt.Println(Hello("benchmark!"))
	}
}
