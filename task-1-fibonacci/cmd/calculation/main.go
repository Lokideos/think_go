package main

import (
	"fmt"

	"github.com/Lokideos/think_go/task-1-fibonacci/pkg/sequencer"
)

func main() {
	max := 20
	result := sequencer.Fibonacci(max)
	fmt.Println(result)
}
