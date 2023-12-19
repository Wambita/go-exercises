package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	filename := os.Args[0]
	for _, char := range filename[2:] {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
