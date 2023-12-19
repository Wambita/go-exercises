package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	// flag to check if upper is present
	upper := false

	// check if upper flag is prsent and remove it from args
	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	// loop through each arg to turn string into int

	for _, arg := range args {
		position := atoi(arg) // pos of letter as arg

		// range should be btwn 1-26
		if position >= 1 && position <= 26 {
			// get letter in corresponding no.
			letter := 'a' + rune(position-1)

			if upper { // to uppercase
				letter -= 'a' - 'A'
			}
			// print letter
			z01.PrintRune(letter)
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func atoi(s string) int {
	result := 0

	// loop through string chars
	for _, num := range s {
		number := int(num - '0')

		// valid num
		if number >= 0 && number <= 9 {
			result = result*10 + number
		} else {
			return 0
		}

	}
	return result
}
