package diffsquares

func SquareOfSum(n int) int {
	sum := 0
	for i := 0; i <= n ; i++ {
		sum += i
	}

	// sum := n * (n + 1) / 2
	return sum * sum
}

func SumOfSquares(n int) int {
	sum := 0

	for i := 0; i <= n ; i++ {
		sum += i * i
	}

	// return n * (n + 1) * (2*n + 1) / 6
	return sum 
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
