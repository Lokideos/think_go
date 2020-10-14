package main

import (
	"fmt"

	"github.com/Lokideos/think_go/task-1-fibonacci/pkg/sequencer"
)

func main() {
	lastNumberIndex := 20
	result := sequencer.Fibonacci(lastNumberIndex)
	fmt.Println(result)
}
