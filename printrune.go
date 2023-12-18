package piscine

import "github.com/01-edu/z01"

func PrintRune() {
	var aRune string = "Ahh"
	var aRunelength int = len(aRune)
	// z01.PrintRune(aRune)
	// z01.PrintRune('\n')
	for i := 0; i <= aRunelength; i++ {
		z01.PrintRune(rune(aRune[1]))
	}
}
