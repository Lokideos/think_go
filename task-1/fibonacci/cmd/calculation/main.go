package main

import (
	"fmt"

	"think.go/task-1/fibonacci/pkg/math"
)

func main() {
	max := 20
	result := math.Fibonacci(max)
	fmt.Println(result)
}
