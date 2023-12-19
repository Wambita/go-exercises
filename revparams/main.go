package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	index := len(args) - 1

	for index >= 0 { // loop through every word

		for _, str := range args[index] { // loop through every char in word
			z01.PrintRune(str)
		}
		index--
		z01.PrintRune('\n')
	}
}
