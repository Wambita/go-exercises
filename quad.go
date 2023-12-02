package piscine

import "github.com/01-edu/z01"

func main() {
	z01.PrintRune('\n')
	x := 5
	// y := 3
	for z := 1; z <= x; z++ {
		if x == 1 && x != 0 {
			z01.PrintRune('o')
		} else if x == 2 && x == x-1 {
			for a := 2; a <= x-1; a++ {
				z01.PrintRune('-')
			}
		} else if x == x {
			z01.PrintRune('o')
		}
	}
}
