package main

import (
	"flag"
	"fmt"

	"github.com/Lokideos/think_go/task-1-fibonacci/pkg/sequencer"
)

const (
	// defaultFibnonacciMax is a default maximum nubmer during search for last number in Fibonacci sequence
	defaultLastNumberIndex = 20
)

func main() {
	var lastNumberIndex int
	flag.IntVar(&lastNumberIndex, "f", defaultLastNumberIndex, "Number to compute in Fibonacci Sequence")
	flag.Parse()

	result := sequencer.Fibonacci(lastNumberIndex)
	fmt.Println(result)
}
