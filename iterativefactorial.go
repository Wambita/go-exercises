package piscine

// iterative factorial

func IterativeFactorial(nb int) int {
	factorial := 1
	if nb < 0 {
		return 0
	} else if nb > 0 && nb < nb+1 {
		for i := 1; i <= nb; i++ {
			factorial = factorial * i
		}
	}
	return factorial
}
