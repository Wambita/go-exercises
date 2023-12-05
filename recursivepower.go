package main

import "fmt"

// recursive power

func RecursivePower(nb int, power int) int {
	if power == 0 {
		return 1
	} else if power < 0 {
		return 0 // returns 0 for # greater than 24 and less than 0
	} else {
		return nb * RecursivePower(nb, power-1) // calculate power using the n * n recursively
	}
}

func main() {
	result3 := RecursivePower(2, 3)
	fmt.Printf("result: %v", result3)
}
