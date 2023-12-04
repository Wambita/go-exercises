package main

import (
	"fmt"
)

// iterative factorial

func IterativeFactorial(nb int) int {
	factorial := 1

	for i := 1; i <= nb; i++ {
		factorial = factorial * i
	}
	return factorial
}

func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}
