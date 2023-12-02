package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	if x < 0 || y < 0 {
		z01.PrintRune('\n')
	} else {
		// top width
		i := 0
		if i == 0 {
			z01.PrintRune('o')
			for w := 0; w < x-2; w++ {
				z01.PrintRune('-')
			}
			last := x
			if last == x {
				z01.PrintRune('o')
			}
			z01.PrintRune('\n')
		}
		// widths
		for j := 0; j < y-2; j++ {
			z01.PrintRune('|')
			// for space between the heights
			for s := 0; s < x-2; s++ {
				z01.PrintRune(' ')
			}

			z01.PrintRune('|')
			z01.PrintRune('\n')
		} // bottom width
		t := 0
		if t == 0 {
			z01.PrintRune('o')
			for w := 0; w < x-2; w++ {
				z01.PrintRune('-')
			}
			last := x
			if last == x {
				z01.PrintRune('o')
			}
			z01.PrintRune('\n')
		}
	}
	z01.PrintRune('\n')
}

func main() {
	QuadA(1, 5)
}
