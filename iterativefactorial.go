package piscine

// iterative factorial

func IterativeFactorial(nb int) int {
	factorial := 1
	for i := 1; i <= nb; i++ {
		factorial = factorial * i
	}
	return factorial
}
