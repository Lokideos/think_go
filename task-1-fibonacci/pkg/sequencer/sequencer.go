package sequencer

// Fibonacci function calculates last fibonacci number in fibonacci sequence given the provided 'max' variable via user input
// The result have to be less than 'max' variable
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
