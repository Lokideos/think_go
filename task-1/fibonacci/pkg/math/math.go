package math

// Fibonacci function calculates last fibonacci number in fibonacci sequence given the provided 'max' variable via user input
// The result have to be less than 'max' variable
func Fibonacci(max int) int {
	fibArr := []int{1, 1}

	for nextElement(fibArr) < max {
		fibArr = append(fibArr, nextElement(fibArr))
	}

	return fibArr[len(fibArr)-1]
}

func nextElement(arr []int) int {
	return (arr[len(arr)-1] + arr[len(arr)-2])
}
