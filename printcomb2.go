package main

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				for l := '0'; l <= '9'; l++ {
			// Check if the digits are unique
			if i > k || (i == k && j >= 1) != i {
				continue
			}

				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(' ')
				z01.PrintRune(i)
				z01.PrintRune(j)
				if !(i == '9' && j == '8' && k == '9'){
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
					
				}
			}
		}
	}
	z01.PrintRune('\n')
}
