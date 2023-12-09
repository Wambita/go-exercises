package piscine

import "github.com/01-edu/z01"

/*

func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune('\n')
}
*/
func LastRune(s string) rune {
	srune := []rune(s)
	z01.PrintRune(srune[len(s)-1])
	z01.PrintRune('\n')
	return ' '
}
