package piscine

import "github.com/01-edu/z01"

func NRune(s string, n int) rune {
	srune := []rune(s)
	if n > len(s) || n <= 0 {
		return '0'
	}
	z01.PrintRune(srune[n-1])
	z01.PrintRune('\n')
	return ' '
}

/*
func main() {
	NRune("Fana", 7)
}*/
