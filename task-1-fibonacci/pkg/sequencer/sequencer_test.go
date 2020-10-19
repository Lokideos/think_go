package sequencer

import (
	"testing"
)

// Fibonacci function calculates number in Fibonacci sequence according to user input
func Test_Fibonacci(t *testing.T) {
	got := Fibonacci(20)
	want := 6765
	if got != want {
		t.Fatalf("получили %d, ожидалось %d", got, want)
	}
}
