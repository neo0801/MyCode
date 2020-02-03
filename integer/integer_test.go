package integer

import (
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 3)
	want := 5

	if sum != want {
		t.Errorf("got %d, want %d", sum, want)
	}
}
