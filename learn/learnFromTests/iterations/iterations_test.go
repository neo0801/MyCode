package iterations

import (
	"testing"
)

func TestIteration(t *testing.T) {
	got := LoopFor("a", 3)
	want := "aaa"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
