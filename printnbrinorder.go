package piscine

import (
	"github.com/01-edu/z01"
)

/*
func main() {
	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)
}*/

func PrintNbrInOrder(n int) {
	// array to store count of each digit
	count := make([]int, 10)

	// count occurence of each digit
	if n == 0 {
		z01.PrintRune('0')
	}

	for n > 0 {
		number := n % 10
		count[number]++
		n /= 10

	}
	// print digits in ascending order
	for i := 0; i < 10; i++ {
		for j := 0; j < count[i]; j++ {
			tuRune := '0' + rune(i)
			z01.PrintRune(tuRune)
		}
	}
}
