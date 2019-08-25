// Package diffsquares implement some functions
package diffsquares

// SquareOfSum returns squares of sum natural number
func SquareOfSum(n int) int {
	x := n * (n + 1) / 2
	return x * x
}

// SumOfSquares return sum of squares natural number
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference return difference between SquareOfSum and SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
