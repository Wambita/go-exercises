package main

// iterative power

func IterativePower(nb int, power int) int {
	powers := 1
	if nb == 1 || nb == 0 {
		return 1
	} else if nb > 1 && nb <= 23 {
		for i := 0; i < power; i++ {
			powers = powers * nb // calculate power using the formula n * n iteratively
		}
	} else {
		return 0 // returns 0 for # greater than 24 and less than 0
	}
	return powers
}
