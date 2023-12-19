package piscine

import (
	"github.com/01-edu/z01"
)

/*
func main() {
	PrintNbrBase(125, "0123456789")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "01")
	z01.PrintRune('\n')
	PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	PrintNbrBase(125, "aa")
	z01.PrintRune('\n')
}
*/
func Indexs(s string, toFind string) int {
	// returns the index of the first instance of substring return -1 if not present

	// calculate lengths of input strings
	toFindLen := len(toFind)

	// iterate through each possible starting index of a substring in s
	for i := 0; i <= len(s)-len(toFind); i++ {
		// variable to track if a match is found for current index
		match := true

		// iterate through each char to find a substring
		for j := 0; j < toFindLen; j++ {
			if s[i+j] != toFind[j] { // check if theres a mismatch
				match = false
				break
			}
		}
		// if match is true return i
		if match {
			return i
		}
	}
	// return -1 if no match was found
	return -1
}

func PrintBaseNum(n int, base string) {
	// calc modulo of base to get correct pos of char
	modulo := len(base)
	i := 0

	// if num = 0, print 0 and return
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	// loop to find pos of the char in the base for the  given modulo
	for j := 1; j <= n%modulo; j++ {
		i++
	}

	// loop to find pos of char in base for the given neg modulo
	for j := -1; j >= n%modulo; j-- {
		i++
	}
	// recursion to print num in specified base
	if n/modulo != 0 {
		PrintBaseNum(n/modulo, base)
	}
	// Print char corresponding to the calculated pos in the base
	z01.PrintRune(rune(base[i]))
	return
}

func PrintNbrBase(nbr int, base string) {
	// check if base is invalid (less than 2 , contains + or -)
	if len(base) < 2 || Indexs(base, "+") >= 0 || Indexs(base, "-") >= 0 {
		// Print NV for invalid
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	// check for duplicates in the base
	for i, char := range base {
		for j, chars := range base {
			if i != j && char == chars {
				// print nv for invalid duplicate chars
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}
	// negative nums
	if nbr < 0 {
		z01.PrintRune('-')
	}
	PrintBaseNum(nbr, base)
}
