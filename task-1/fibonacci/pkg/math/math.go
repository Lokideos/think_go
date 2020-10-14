package math

// Fibonacci function calculates last fibonacci number in fibonacci sequence given the provided 'max' variable via user input
// The result have to be less than 'max' variable
func Fibonacci(max int) int {
	fibonacciSequence := []int{1, 1}

	for nextElement(fibonacciSequence) < max {
		fibonacciSequence = append(fibonacciSequence, nextElement(fibonacciSequence))
	}

	return fibonacciSequence[len(fibonacciSequence)-1]
}

func nextElement(arr []int) int {
	return (arr[len(arr)-1] + arr[len(arr)-2])
}
