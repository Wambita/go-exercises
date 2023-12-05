package piscine

// iterative factorial

func IterativeFactorial(nb int) int {
	factorial := 1
	if nb == 1 || nb == 0 {
		return 1
	} else if nb > 1 && nb <= 23 {
		for i := 1; i <= nb; i++ {
			factorial = factorial * i
		}
	} else {
		return 0 // returns 0 for # greater than 24 and less than 0
	}
	return factorial
}
