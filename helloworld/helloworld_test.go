package main

import (
	"testing"
)

// TestHello tests Hello function
func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world!"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
