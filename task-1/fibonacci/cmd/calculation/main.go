package main

import (
	"fmt"

	"../../pkg/math"
)

func main() {
	max := 20
	result := math.Fibonacci(max)
	fmt.Println(result)
}
