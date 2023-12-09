package piscine

import "github.com/01-edu/z01"

func StrRev(s string) {
	srune := []rune(s)                 // change string to rune
	for i := len(s) - 1; i >= 0; i-- { // iterate through loop from last index item to the first
		z01.PrintRune(srune[i]) // print rune element at that index
	}
}

/*
func main() {
	StrRev("fana")
}
*/
