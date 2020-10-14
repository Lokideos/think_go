package sequencer

// Fibonacci function calculates last number in Fibonacci sequence according to user input
func Fibonacci(lastNumberIndex int) int {
	sequence := []int{1, 1}

	for len(sequence) < lastNumberIndex {
		sequence = append(sequence, nextFibonacciElement(sequence))
	}

	return sequence[len(sequence)-1]
}

func nextFibonacciElement(arr []int) int {
	return (arr[len(arr)-1] + arr[len(arr)-2])
}
