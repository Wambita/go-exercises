package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	// get 4 numbers loop
	for i := '0'; i <= '9'; i++ { // first digit first number
		for j := '0'; j <= '9'; j++ { // second digit first num
			for k := '0'; k <= '9'; k++ { // first digit second num
				for l := '0'; l <= '9'; l++ { // second digit second num

					if i != k && j != l { // ensures the first combination is not equal to the second combination
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(k)
						z01.PrintRune(l)

						if i != '9' && j != '9' && k != '9' && l != '8' { // no comma on the last number
							z01.PrintRune(' ')
							z01.PrintRune(',')

						}

					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
