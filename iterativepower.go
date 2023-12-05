package main

// iterative power

func IterativePower(nb int, power int) int {
	powers := 1
	if nb == 1 || nb == 0 {
		return 1
	} else if nb > 1 && nb <= 23 {
		for i := 0; i < power; i++ {
			powers = powers * nb
		}
	} else {
		return 0
	}
	return powers
}
