package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	for _, char := range args { // loop through every word
		for _, str := range char { // loop through every char in word
			z01.PrintRune(str)
		}

		z01.PrintRune('\n')
	}
}
