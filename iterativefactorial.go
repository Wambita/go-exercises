package piscine

// iterative factorial

func IterativeFactorial(nb int) int {
	factorial := 1
	if nb == 0 || nb < 0 {
		return 0
	} else if nb > 0 && nb <= 23 {
		for i := 1; i <= nb; i++ {
			factorial = factorial * i
		}
	} else {
		return 0
	}
	return factorial
}
