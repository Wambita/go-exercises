package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			// Check if the digits are unique
			if i != j || j != i {

				z01.PrintRune(i)
				z01.PrintRune(j)

				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}

/*func main() {
	PrintComb()
}*/
