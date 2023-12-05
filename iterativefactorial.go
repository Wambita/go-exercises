package piscine

// iterative factorial

func IterativeFactorial(nb int) int {
	factorial := 1
	if nb > 24 && nb < 0 {
		return 1
	} else {
		for i := 1; i <= nb; i++ {
			factorial = factorial * i
		}
	}
	return factorial
}
