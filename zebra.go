package piscine

import "github.com/01-edu/z01"

func main() {
	var aRune string = "Ahh"
	var aRunelength int = len(aRune)
	z01.PrintRune(rune(aRunelength))
}

/*z01.PrintRune('\n')
for i := 0; i <= aRunelength; i++ {
	for s := aRune[i]; s <= aRune[aRunelength-1]; s++ {
		if aRune[i] == 'z' {
			z01.PrintRune(rune(aRune[0]))
		} else {
			z01.PrintRune('z')
			z01.PrintRune('\n')
		}
	}
}
*/
