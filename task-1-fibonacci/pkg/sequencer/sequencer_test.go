package sequencer

import (
	"testing"
)

// Fibonacci function calculates number fibonacci according to user inpu
func Test_Fibonacci(t *testing.T) {
	got := Fibonacci(20)
	want := 6765
	if got != want {
		t.Fatalf("получили %d, ожидалось %d", got, want)
	}
}
